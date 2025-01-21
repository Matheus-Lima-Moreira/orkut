$('#signup-form').on('submit', Signup);

function Signup(e) {
  e.preventDefault();

  if ($('#password').val() !== $('#confirmPassword').val()) {
    Swal.fire(
      'Error',
      'The password and confirm password doesn\'t match!',
      'error'
    )
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
    Swal.fire(
      'Success',
      'Signup successful',
      'success'
    ).then(() => {
      $.ajax({
        url: "/login",
        method: "POST",
        data: {
          email: $('#email').val(),
          password: $('#password').val()
        }
      }).done(() => {
        window.location = "/home";
      }).fail(() => {
        Swal.fire(
          'Error',
          'Failed to login. Please try again',
          'error'
        )
      })
    });
  }).fail((_) => {
    Swal.fire(
      'Error',
      'Failed to signup. Please try again',
      'error'
    )
  });
}