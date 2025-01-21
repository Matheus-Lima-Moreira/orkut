$('#unfollow').on('click', Unfollow)
$('#follow').on('click', Follow)

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