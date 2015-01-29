define(["jquery"], function($){
	var form = $("#signup-form");
	form.on("submit", function(e){
		e.preventDefault();
		var passwd="";
		var valid = true;
		$(".form-group").removeClass("has-error");
		$("small").addClass("hidden");
		$("input", this).each(function(){
			this.checkValidity();
			if(!this.validity.valid){
				valid = false;
				$(this).closest(".form-group").addClass("has-error");
				$(this).siblings("small").removeClass("hidden");
			}
			if(this.name === "password") passwd=this.value;
			if(this.name === "password_confirm" && this.value !== passwd){
				valid = false;
				$(this).closest(".form-group").addClass("has-error");
				$(this).siblings("small").removeClass("hidden");
			}
		});
		if(valid) form[0].submit();
	});
});
