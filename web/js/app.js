document.getElementById('swagger-btn').addEventListener('click', function () {
    const currentUrl = window.location.origin;
    window.location.href = currentUrl + "/swagger/index.html";
});
