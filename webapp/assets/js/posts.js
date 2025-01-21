$(document).on('submit', '#new-post', CreatePost)
$(document).on('click', '.like-post', LikePost)
$(document).on('click', '.dislike-post', DislikePost)
$(document).on('click', '.delete-post', DeletePost)
$(document).on('click', '#update-post', UpdatePost)

function CreatePost(e) {
  e.preventDefault();

  $.ajax({
    url: "/posts",
    method: "POST",
    data: {
      title: $('#title').val(),
      content: $('#content').val()
    }
  }).done(() => {
    window.location.reload();
  }).fail(() => {
    Swal.fire(
      'Error',
      'There was an error creating the post - please try again later',
      'error'
    )
  });
}

function LikePost(e) {
  e.preventDefault();

  const clickedElement = $(e.target);
  const postId = clickedElement.closest('div').data('post-id');

  clickedElement.prop('disabled', true);
  $.ajax({
    url: `/posts/${postId}/like`,
    method: "POST"
  }).done(() => {
    const countLikes = $(e.target).next('span')
    countLikes.text(parseInt(countLikes.text()) + 1);

    clickedElement.addClass('dislike-post');
    clickedElement.addClass('text-danger');
    clickedElement.removeClass('like-post');
  }).fail(() => {
    Swal.fire(
      'Error',
      'There was an error liking the post - please try again later',
      'error'
    )
  }).always(() => {
    clickedElement.prop('disabled', false);
  });
}

function DislikePost(e) {
  e.preventDefault();

  const clickedElement = $(e.target);
  const postId = clickedElement.closest('div').data('post-id');

  clickedElement.prop('disabled', true);
  $.ajax({
    url: `/posts/${postId}/dislike`,
    method: "POST"
  }).done(() => {
    const countLikes = $(e.target).next('span')
    countLikes.text(parseInt(countLikes.text()) - 1);

    clickedElement.addClass('like-post');
    clickedElement.removeClass('text-danger');
    clickedElement.removeClass('dislike-post');
  }).fail(() => {
    Swal.fire(
      'Error',
      'There was an error disliking the post - please try again later',
      'error'
    )
  }).always(() => {
    clickedElement.prop('disabled', false);
  });
}

function UpdatePost(e) {
  e.preventDefault();

  $(this).prop('disabled', true);

  const postId = $(this).data('post-id');

  $.ajax({
    url: `/posts/${postId}`,
    method: "PUT",
    data: {
      title: $('#title').val(),
      content: $('#content').val()
    }
  }).done(() => {
    Swal.fire(
      'Success',
      'Post updated successfully',
      'success'
    ).then(() => {
      window.location = '/home'
    })
  }).fail(() => {
    Swal.fire(
      'Error',
      'There was an error editing the post - please try again later',
      'error'
    )
  }).always(() => {
    $('#update-post').prop('disabled', false)
  });
}

function DeletePost(e) {
  e.preventDefault();

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
      deletePostAjax(e);
    }
  });

  function deletePostAjax(e) {
    const clickedElement = $(e.target);
    const postElement = clickedElement.closest('div');
    const postId = postElement.data('post-id');

    clickedElement.prop('disabled', true);
    $.ajax({
      url: `/posts/${postId}`,
      method: "DELETE"
    }).done(() => {
      postElement.fadeOut("slow", function () {
        postElement.remove();
      });
    }).fail(() => {
      Swal.fire(
        'Error',
        'There was an error deleting the post - please try again later',
        'error'
      )
    });
  }
}