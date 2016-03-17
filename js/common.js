$(document).ready(function(){
	$("time").each(function(i, elm){
		var utc = moment.utc($(elm).attr('datetime'));
		$(elm).html(utc.local().format('LLLL'));
		$(elm).closest("article").show();
	});
});