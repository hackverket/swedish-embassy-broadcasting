$( document ).ready(function() {

  $('#circle').circleProgress({
    value: 0.0,
    size: 80,
    fill: {
      gradient: ["red", "orange"]
    }
  });


	getData();

	setInterval(function(){ getData(); }, 3000);
});

function getData() {
	$.getJSON( "/api/queue/", function( data ) {
		$(".entry1-img").attr("src",data[0].Image);
		$("#entry1-text").text("NOW PLAYING: " + data[0].Title);
		$('.korv').html('');

    if (data.length > 1) {
      var past = data[0].Duration * -1;
      var next = data[1].Duration;
      var progress = past / (past + next);
      $('#circle').circleProgress('value', progress); 
    } else {
      $('#circle').circleProgress('value', 0); 
    }

		_.forEach(data, function (value, key) {
			if(key > 0 && key < 4) {
				$('.korv').append(''+
				'<div class="col-lg-4">' +
          		'<img class="img-circle" src="'+value.Image+'"" alt="' + value.Title+'" width="200" height="200">' +
          		'<h2>' + value.Title+'</h2>' +
        		'</div>')
			}
		});
	});
}
