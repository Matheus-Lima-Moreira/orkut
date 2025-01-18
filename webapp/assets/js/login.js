$('#login-form').on('submit', Login);

function Login(e) {
  e.preventDefault();

  $.ajax({
    url: "/login",
    method: "POST",
    data: {
      email: $('#email').val(),
      password: $('#password').val()
    }
  }).done(() => {
    window.location = "/home";
  }).fail((err) => {
    console.log(err);
    alert("Invalid credentials");
  });
}