$('#new-post').on('submit', createPost);
$(document).on('click', '.like-post', likePost)
$(document).on('click', '.dislike-post', dislikePost)

function createPost(e) {
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
    alert('There was an error creating the post - please try again later');
  });
}

function likePost(e) {
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
    alert('There was an error liking the post - please try again later');
  }).always(() => {
    clickedElement.prop('disabled', false);
  });
}

function dislikePost(e) {
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
    alert('There was an error disliking the post - please try again later');
  }).always(() => {
    clickedElement.prop('disabled', false);
  });
}