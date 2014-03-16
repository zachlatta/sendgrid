function addItem(text) {
  docHeight = $(document).innerHeight() * 0.9;
  docWidth = $(document).innerWidth() * 0.9;
  textCol = '#' + Math.floor(Math.random() * 16777215).toString(16);
  textSize = Math.floor((Math.random() * 48) + 12);
  yPos = Math.floor((Math.random() * docHeight));
  xPos = Math.floor((Math.random() * docWidth));
  $('<div class="wow" style="color:' + textCol + ';left:' + xPos + 'px;top:' + yPos + 'px;font-size:' + textSize + 'px;">' + text + '</div>').appendTo("body").fadeIn('fast').delay(4000).fadeOut('fast');
}
