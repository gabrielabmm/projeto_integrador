if (acceptButton && cookieBanner) {
    acceptButton.addEventListener('click', function () {
        fetch("/set-cookie-consent?value=accepted", { method: "POST" });
        cookieBanner.style.display = 'none';
    });
}

if (declineButton && cookieBanner) {
    declineButton.addEventListener('click', function () {
        fetch("/set-cookie-consent?value=declined", { method: "POST" });
        cookieBanner.style.display = 'none';
    });
}
