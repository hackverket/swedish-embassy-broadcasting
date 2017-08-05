$( document ).ready(function() {
	getData();

	setTimeout(function(){ getData(); }, 3000);
});

function getData() {
	$.getJSON( "http://10.44.22.141:4020/api/queue/", function( data ) {
		$(".entry1-img").attr("src",data[0].Image);
		$("#entry1-text").text("NOW PLAYING: " + data[0].Title);
		$('.korv').html('');

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