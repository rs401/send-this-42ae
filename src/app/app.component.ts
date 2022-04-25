import { Component, OnInit } from '@angular/core';
import { Meta, Title } from '@angular/platform-browser';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'send-this';

  constructor(private titleSvc: Title, private metaSvc: Meta) { }

  ngOnInit(): void {
    this.titleSvc.setTitle(this.title);
    this.metaSvc.addTags([
      {name: 'keywords', content: 'send to phone, send file to phone, send file to mobile, send this to phone, send this to mobile'},
      {name: 'description', content: 'Send this file to my phone/mobile, anything that can scan a QR code.'},
      {name: 'robots', content: 'index, follow'}
    ]);
  }
}
