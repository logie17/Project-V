requirejs.config({
	paths: {
		jquery: "vendor/jquery",
		bootstrap: "vendor/bootstrap"
	}
});
require(['jquery', "bootstrap"], function( $ ) {
	    console.log( $ ) // OK
});
require(["app/home"], function( $ ) {
});
