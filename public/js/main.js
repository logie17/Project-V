requirejs.config({
	paths: {
		jquery: "vendor/jquery"
	}
});
require(['jquery'], function( $ ) {
	    console.log( $ ) // OK
});
require(["app/home"], function( $ ) {
});
