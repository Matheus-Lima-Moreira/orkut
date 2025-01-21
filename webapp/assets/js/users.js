$('#unfollow').on('click', Unfollow)
$('#follow').on('click', Follow)
$('#edit-profile').on('submit', Edit)
$('#update-password').on('submit', UpdatePassword)
$('#delete-account').on('click', Delete)

function Follow() {
  const userId = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${userId}/follow`,
    method: "POST"
  }).done(() => {
    window.location.reload();
  }).fail(() => {
    Swal.fire(
      'Error',
      'Failed to follow user - please try again later',
      'error'
    )
    $('#follow').prop('disabled', false)
  });
}

function Unfollow() {
  const userId = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${userId}/unfollow`,
    method: "POST"
  }).done(() => {
    window.location.reload();
  }).fail(() => {
    Swal.fire(
      'Error',
      'Failed to unfollow user - please try again later',
      'error'
    )
    $('#unfollow').prop('disabled', false)
  });
}

function Edit(e) {
  e.preventDefault();

  $.ajax({
    url: `/edit-profile`,
    method: "PUT",
    data: {
      name: $('#name').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
    }
  }).done(() => {
    Swal.fire(
      'Success',
      'Profile updated successfully',
      'success'
    ).then(() => {
      window.location = "/profile"
    });
  }).fail(() => {
    Swal.fire(
      'Error',
      'Failed to update user name - please try again later',
      'error'
    )
  });
}

function UpdatePassword(e) {
  e.preventDefault();

  if ($('#newPassword').val() != $('#confirmPassword').val()) {
    Swal.fire(
      'Error',
      'The new password and confirm password doesn\'t match!',
      'error'
    )
    return;
  }

  $.ajax({
    url: `/update-password`,
    method: "PUT",
    data: {
      current_password: $('#currentPassword').val(),
      new_password: $('#newPassword').val()
    }
  }).done(() => {
    Swal.fire(
      'Success',
      'Password updated successfully',
      'success'
    ).then(() => {
      window.location = "/profile"
    });
  }).fail(() => {
    Swal.fire(
      'Error',
      'Failed to update password - please try again later',
      'error'
    )
  });
}

function Delete() {
  Swal.fire({
    title: 'Are you sure?',
    text: "You won't be able to revert this!",
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#655CC9',
    cancelButtonColor: '#d33',
    confirmButtonText: 'Yes, delete it!'
  }).then((result) => {
    if (result.isConfirmed) {
      $.ajax({
        url: `/delete-account`,
        method: "DELETE"
      }).done(() => {
        Swal.fire(
          'Success!',
          'Your account has been deleted.',
          'success'
        ).then(() => {
          window.location = "/logout"
        });
      }).fail(() => {
        Swal.fire(
          'Error',
          'Failed to delete account - please try again later',
          'error'
        )
      });
    }
  });
}