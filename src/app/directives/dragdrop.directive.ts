import { Directive, EventEmitter, HostListener, Output } from '@angular/core';

@Directive({
  selector: '[appDragdrop]'
})
export class DragdropDirective {
  @Output() fileDropped = new EventEmitter<any>();

  constructor() { }

  @HostListener('drop', ['$event']) public ondrop(event: any) {
    event.preventDefault();
    event.stopPropagation();
    let files = event.dataTransfer.files;
    if (files.length > 0) {
      this.fileDropped.emit(files);
    }
  }

}
