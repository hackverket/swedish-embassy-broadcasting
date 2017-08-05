$( document ).ready(function() {
	getData();
});

function getData() {
	$.getJSON( "http://10.44.22.141:4020/api/queue/", function( data ) {
		console.log(data);
	});
}