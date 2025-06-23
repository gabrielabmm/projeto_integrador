sacForm.addEventListener('submit', function(event) {
    event.preventDefault();

    emailErrorDiv.style.display = 'none';
    emailInput.classList.remove('invalid');
    sacSuccessMessageDiv.style.display = 'none';

    const emailValue = emailInput.value.trim();
    const messageValue = messageTextarea.value.trim();

    fetch('/api/sac', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email: emailValue, mensagem: messageValue })
    })
    .then(response => {
        if (!response.ok) {
            return response.text().then(text => { throw new Error(text); });
        }
        return response.text();
    })
    .then(data => {
        labelEmail.style.display = 'none';
        emailInput.style.display = 'none';
        labelMensagem.style.display = 'none';
        messageTextarea.style.display = 'none';
        submitButton.style.display = 'none';

        sacSuccessMessageDiv.textContent = data;
        sacSuccessMessageDiv.style.display = 'block';
    })
    .catch(error => {
        emailErrorDiv.textContent = error.message;
        emailErrorDiv.style.display = 'block';
        emailInput.classList.add('invalid');
    });
});

