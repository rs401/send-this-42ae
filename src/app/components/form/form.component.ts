import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component, HostListener, OnInit, ViewChild } from '@angular/core';
import { NgbAlert } from '@ng-bootstrap/ng-bootstrap';
import { debounceTime, map, Subject } from 'rxjs';
import { environment } from 'src/environments/environment';


interface GeturlsResponse {
  upload_url: string;
  download_name: string;
}

@Component({
  selector: 'app-form',
  templateUrl: './form.component.html',
  styleUrls: ['./form.component.css']
})
export class FormComponent implements OnInit {
  private _error = new Subject<string>();
  errMsg: string = '';
  file!: File;
  backend: string;
  geturlsResponse!: GeturlsResponse;
  showForm: boolean = true;
  showSpinner: boolean = false;
  showQr: boolean = false;
  qrURL: string = "";

  @ViewChild('fileSizeLimitAlert', {static: false}) fileSizeLimitAlert!: NgbAlert;

  constructor(private http: HttpClient) {
    this.backend = environment.backend;
  }

  ngOnInit(): void {
    this._error.subscribe(msg => this.errMsg = msg);
    this._error.pipe(debounceTime(5000)).subscribe(() => {
      if(this.fileSizeLimitAlert) {
        this.fileSizeLimitAlert.close();
      }
    });
  }

  displayError(msg: string) {
    this._error.next(msg);
  }

  fileChangeHandler(event: any): void {
    // Start spinner
    this.showSpinner = true;
    this.showForm = false;
    let tmpFile: File = (event.target.files as FileList)[0];
    if (!tmpFile) {
      tmpFile = (event.dataTransfer.files as FileList)[0];
    }
    // Limit file size
    if(tmpFile.size/1024/1024 > 10) {
      // Stop spinner and show form again
      this.showSpinner = false;
      this.showForm = true;
      // set error message and return
      this.displayError('ERROR: File size must not exceed 10Mb.');
      return
    }
    
    const formData = new FormData();
    formData.append('filetype', tmpFile.type)
    this.file = tmpFile;
    console.log('form event: ' + tmpFile.name);
    const getIt = this.http.post(this.backend + 'geturls', formData).pipe(
      map(data => {
        this.geturlsResponse = data as GeturlsResponse;
        this.sendFile();
      })
    );
    getIt.subscribe({
      next: () => {
        this.genQR('temp/' + this.geturlsResponse.download_name);
      }
    });
  }

  @HostListener("drop", ["$event"]) onDrop(event: any) {
    console.log('hostlistener drop event: ' + JSON.stringify(event));
    event.preventDefault();
    event.stopPropagation();
    this.fileChangeHandler(event);
  }

  sendFile() {
    this.http.put(
      this.geturlsResponse.upload_url,
      this.file,
      { headers: new HttpHeaders({
        'Content-Type':  this.file.type,
      }) },).subscribe({
      next: (val) => {
        console.log('sendFile next: (val) => ' + JSON.stringify(val));
      },
      error: (err) => {
        console.log('sendFile error: ' + JSON.stringify(err));
      },
    });
  }

  genQR(name: string) {
    // Synthesize download url
    let downUrl: string = environment.backend + 'downurl/' + name;
    // Gen qr image
    this.qrURL = downUrl;
    // Stop spinner
    this.showSpinner = false;
    // Display qr image
    this.showQr = true;
  }

}
