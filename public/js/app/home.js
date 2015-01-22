define(["jquery"], function($){
	$("body").on("blur click dblclick focus", function(e){
		console.log(e.type, e.target);
	});	
});
