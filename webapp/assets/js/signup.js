$('#signup-form').on('submit', Signup);

function Signup(e) {
  e.preventDefault();

  if ($('#password').val() !== $('#confirmPassword').val()) {
    alert('Please enter your password again and try again.');
  }

  $.ajax({
    url: "/signup",
    method: "POST",
    data: {
      name: $('#name').val(),
      nickname: $('#nickname').val(),
      email: $('#email').val(),
      password: $('#password').val()
    }
  });
}