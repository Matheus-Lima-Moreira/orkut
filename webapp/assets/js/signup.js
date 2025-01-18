$('#signup-form').on('submit', Login);

function Login(e) {
  e.preventDefault();

  if ($('#password').val() !== $('#confirmPassword').val()) {
    alert('Please enter your password again and try again.');
    return;
  }

  $.ajax({
    url: "/signup",
    method: "POST",
    data: {
      name: $('#name').val(),
      nick: $('#nick').val(),
      email: $('#email').val(),
      password: $('#password').val()
    }
  }).done(() => {
    alert("Signup successful")
  }).fail((err) => {
    console.log(err);
    alert("Failed to signup. Please try again");
  });
}