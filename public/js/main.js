requirejs.config({
	paths: {
		jquery: "vendor/jquery",
		swig: "vendor/swig",
		bootstrap: "vendor/bootstrap"
	}
});
require(['jquery', "bootstrap"], function( $ ) {
	    console.log( $ ) // OK
});
require(["app/home"], function( ) {
});
