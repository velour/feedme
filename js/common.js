var feedme = {
	collapseAll: function () {
		$(".winbody").hide();
		$(".wintag").removeClass("expanded");
	},

	expandAll: function () {
		$(".winbody").show();
		$(".wintag").addClass("expanded");
	},
}

$(document).ready(function(){
	$("time").each(function(i, elm){
		var utc = moment.utc($(elm).attr('datetime'));
		$(elm).html(utc.local().format('LLLL'));
		$(elm).closest("article").show();
	});

	$(".wintag .box").click(function(event){
		var win = $(event.target).closest(".win");
		var tag = win.find(".wintag");
		var body = win.find(".winbody");
		if (body.css("display") === "none") {
			body.show();
			tag.addClass("expanded");
		} else {
			body.hide();
			tag.removeClass("expanded");
		}
	})
});