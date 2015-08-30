function login () {
// 	var loginForm = document.forms["loginForm"];	
	var values = {};
	$.each($('#loginForm').serializeArray(), function(i, field) {
    	values[field.name] = field.value;
	});
	alert(values['email']);
	alert(values['password']);
	
	$.post( "/login", $('#loginForm').serialize())
		.done(function (data){
			alert("Returned data:" + data);
		});
	
	return false;
	}