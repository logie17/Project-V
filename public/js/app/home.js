define(["jquery"], function($){
	$("body").on("blur click dblclick focus", function(e){
		console.log(e.type, e.target);
	});
	$(".right-pane .btn").popover({
		content: "<ul class='list-unstyled'><li>Chat</li><li>Video</li><li>tmux</li></ul>",
		html: true,
		template: '<div class="popover" role="tooltip"><div class="arrow"></div><div class="popover-content"></div></div>'
	});
});
