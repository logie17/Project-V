define(["jquery"], function($){
	$(".form-group").removeClass("has-error");
	$("small").addClass("hidden");
	var form = $("form");
	var valid = true;
	form.on("submit", function(e){
		e.preventDefault();
		$("input", this).each(function(){
			this.checkValidity();
			if(!this.validity.valid){
				valid = false;
				$(this).closest(".form-group").addClass("has-error");
				$(this).siblings("small").removeClass("hidden");
			}
		});
		if(valid) form.submit();
	});
});
