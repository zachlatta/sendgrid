var socket = new WebSocket('ws://hack.zachlatta.com/socket');

socket.onopen = function () {
  console.log('Socket opened.');
}

socket.onmessage = function (msg) {
  console.log('Message received: ' + msg);
  docHeight = $(document).innerHeight() * 0.9;
  docWidth = $(document).innerWidth() * 0.9;
  textCol = '#' + Math.floor(Math.random() * 16777215).toString(16);
  textSize = Math.floor((Math.random() * 48) + 12);
  yPos = Math.floor((Math.random() * docHeight));
  xPos = Math.floor((Math.random() * docWidth));
  $('<div class="wow" style="color:' + textCol + ';left:' + xPos + 'px;top:' + yPos + 'px;font-size:' + textSize + 'px;">' + msg + '</div>').appendTo("body").fadeIn('fast').delay(4000).fadeOut('fast');
}
