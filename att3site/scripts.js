document.addEventListener('DOMContentLoaded', function() {
    // Obter referências aos elementos HTML
    const emailInput = document.getElementById('email');
    const emailError = document.getElementById('emailError');
    const enviarCodigoBtn = document.getElementById('enviarCodigoBtn');

    // Função para validar o formato do e-mail usando uma expressão regular
    function isValidEmail(email) {
        // Expressão regular para validação de e-mail (básica, mas eficaz)
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
    }

    // Função para exibir a mensagem de erro
    function showEmailError(message) {
        emailInput.classList.add('invalid'); // Adiciona a classe 'invalid' para estilizar
        emailError.textContent = message;    // Define o texto da mensagem de erro
        emailError.style.display = 'block';  // Torna a mensagem visível
    }

    // Função para esconder a mensagem de erro
    function hideEmailError() {
        emailInput.classList.remove('invalid'); // Remove a classe 'invalid'
        emailError.textContent = '';           // Limpa o texto da mensagem
        emailError.style.display = 'none';     // Esconde a mensagem
    }

    // Adicionar um listener para o evento 'input' no campo de e-mail
    // Isso fará a validação enquanto o usuário digita
    emailInput.addEventListener('input', function() {
        const emailValue = emailInput.value.trim(); // Pega o valor e remove espaços em branco

        if (emailValue === '') {
            hideEmailError(); // Esconde o erro se o campo estiver vazio (apenas para feedback em tempo real)
        } else if (!isValidEmail(emailValue)) {
            showEmailError('Por favor, insira um e-mail válido (ex: seu.email@exemplo.com).');
        } else {
            hideEmailError(); // Esconde o erro se o e-mail estiver válido
        }
    });

    // Adicionar um listener para o evento 'click' no botão "Enviar código"
    enviarCodigoBtn.addEventListener('click', function(event) {
        // Impedir o envio padrão do formulário (para que o JavaScript possa validar)
        event.preventDefault();

        const emailValue = emailInput.value.trim();

        if (emailValue === '') {
            showEmailError('O campo de e-mail é obrigatório.');
        } else if (!isValidEmail(emailValue)) {
            showEmailError('Por favor, insira um e-mail válido (ex: seu.email@exemplo.com).');
        } else {
             // Se o e-mail for válido:
            hideEmailError(); // Esconde qualquer erro existente
            hideEmailSuccess(); // Esconde a mensagem de sucesso (pois vamos redirecionar)

            // ** AQUI ESTÁ A MUDANÇA PRINCIPAL: REDIRECIONAR PARA codigoemail.html **
            window.location.href = 'codigoemail.html';

            // Opcional: Se você quiser passar o e-mail para a próxima página,
            // pode usar sessionStorage ou parâmetros de URL (ex: 'codigoemail.html?email=' + encodeURIComponent(emailValue))
            // Para um caso de uso real, você enviaria o e-mail para o servidor AQUI para gerar o código.
            // Por exemplo:
            /*
            fetch('/api/enviar-codigo', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ email: emailValue })
            })
            .then(response => {
                if (response.ok) {
                    window.location.href = 'codigoemail.html'; // Redireciona após sucesso do servidor
                } else {
                    // Tratar erro do servidor
                    showEmailError('Erro ao enviar código. Tente novamente.');
                }
            })
            .catch(error => {
                console.error('Erro na requisição:', error);
                showEmailError('Erro de conexão. Verifique sua internet.');
            });
            */
        }
    });
});
