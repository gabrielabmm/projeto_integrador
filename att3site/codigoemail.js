document.addEventListener('DOMContentLoaded', function() {
    // Obter referências aos elementos HTML para a página de CÓDIGO
    const codigoInput = document.getElementById('codigoInput');
    const codigoError = document.getElementById('codigoError');
    const codigoSuccessMessage = document.getElementById('codigoSuccessMessage');
    const verificarCodigoBtn = document.getElementById('verificarCodigoBtn');
    
    // Referências para os elementos de reenvio
    const reenviarCodigoBtn = document.getElementById('reenviarCodigoBtn');
    const countdownSpan = document.getElementById('countdown');
    const reenviarInfoSpan = document.getElementById('reenviarInfo');

    let countdownTime = 60; // Tempo inicial do contador em segundos
    let countdownInterval; // Variável para armazenar o ID do intervalo

    // --- Funções Auxiliares de Validação ---
    function isValidCodeFormat(code) {
        // Assume um código de 6 dígitos numéricos.
        // Se o código puder ter letras ou outro formato, ajuste a regex.
        const codeRegex = /^\d{6}$/; // Ex: Apenas 6 dígitos
        return codeRegex.test(code);
    }

    function showCodigoError(message) {
        codigoInput.classList.add('invalid');
        codigoError.textContent = message;
        codigoError.style.display = 'block';
        codigoSuccessMessage.style.display = 'none'; // Esconde a mensagem de sucesso
    }

    function hideCodigoError() {
        codigoInput.classList.remove('invalid');
        codigoError.textContent = '';
        codigoError.style.display = 'none';
    }

    function showCodigoSuccess() {
        codigoSuccessMessage.style.display = 'block'; // Torna a mensagem de sucesso visível
        hideCodigoError(); // Garante que a mensagem de erro esteja oculta
    }

    function hideCodigoSuccess() {
        codigoSuccessMessage.style.display = 'none';
    }

    // --- Funções para o Reenviar Código ---
    function startCountdown() {
        reenviarCodigoBtn.disabled = true; // Desabilita o botão
        // Altera o texto do botão para mostrar o contador
        reenviarCodigoBtn.innerHTML = `Reenviar Código (<span id="countdown">${countdownTime}</span>s)`;
        reenviarInfoSpan.textContent = 'Aguarde antes de reenviar.'; // Mensagem informativa
        
        // Limpa qualquer intervalo anterior para evitar duplicação
        if (countdownInterval) {
            clearInterval(countdownInterval);
        }

        countdownInterval = setInterval(() => {
            countdownTime--;
            reenviarCodigoBtn.querySelector('#countdown').textContent = countdownTime; // Atualiza o span interno

            if (countdownTime <= 0) {
                clearInterval(countdownInterval); // Para o contador
                reenviarCodigoBtn.disabled = false; // Habilita o botão
                reenviarCodigoBtn.textContent = 'Reenviar Código'; // Volta o texto normal do botão
                reenviarInfoSpan.textContent = ''; // Limpa a mensagem informativa
                countdownTime = 60; // Reseta o tempo para a próxima vez
            }
        }, 1000); // Atualiza a cada 1 segundo
    }

    // Inicia o contador assim que a página carrega.
    // Isso simula que um código acabou de ser enviado da página anterior.
    startCountdown(); 

    // --- Listeners de Eventos ---

    // Listener para o evento 'input' no campo do código
    codigoInput.addEventListener('input', function() {
        const codigoValue = codigoInput.value.trim();

        if (codigoValue === '') {
            hideCodigoError();
            hideCodigoSuccess();
        } else if (!isValidCodeFormat(codigoValue)) {
            showCodigoError('O código deve ter 6 dígitos.');
            hideCodigoSuccess();
        } else {
            hideCodigoError(); // Esconde o erro se o formato estiver ok
        }
    });

    // Listener para o evento 'click' no botão "Verificar Código"
    verificarCodigoBtn.addEventListener('click', function(event) {
        event.preventDefault(); // Impede o envio padrão do formulário

        const codigoValue = codigoInput.value.trim();

        if (codigoValue === '') {
            showCodigoError('Por favor, insira o código.');
            hideCodigoSuccess();
        } else if (!isValidCodeFormat(codigoValue)) {
            showCodigoError('O código deve ter 6 dígitos.');
            hideCodigoSuccess();
        } else {
            // Se o formato do código for válido no frontend,
            // AQUI VOCÊ ENVIARIA O CÓDIGO PARA O SEU BACKEND PARA VERIFICAÇÃO.
            // No frontend, você não pode "saber" se o código é válido ou não,
            // a menos que o backend responda a você.

            // --- SIMULAÇÃO DE VALIDAÇÃO DE BACKEND (APENAS PARA TESTE! REMOVA EM PRODUÇÃO!) ---
            const isCodeCorrectFromServer = (codigoValue === '123456'); 

            if (isCodeCorrectFromServer) {
                hideCodigoError();
                showCodigoSuccess(); // Mostra a mensagem de sucesso no frontend
                
                // Em um sistema real, após a validação do backend, você redirecionaria:
                //alert('Código válido! Redirecionando para a página de nova senha...'); 
                window.location.href = 'redefinicao.html'; 

            } else {
                showCodigoError('Código incorreto. Tente novamente.');
                hideCodigoSuccess();
            }

            // Exemplo real com fetch (requer um backend):
            /*
            fetch('/api/verificar-codigo', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ codigo: codigoValue, email: 'email_do_usuario_aqui' }) // Envie o email também!
            })
            .then(response => response.json())
            .then(data => {
                if (data.isValid) {
                    hideCodigoError();
                    showCodigoSuccess();
                    // window.location.href = 'pagina-de-nova-senha.html';
                } else {
                    showCodigoError('Código incorreto. Tente novamente.');
                }
            })
            .catch(error => {
                console.error('Erro na requisição:', error);
                showCodigoError('Erro de conexão. Tente novamente mais tarde.');
            });
            */
        }
    });

    // Listener para o botão "Reenviar Código"
    reenviarCodigoBtn.addEventListener('click', function() {
        // Você enviaria uma requisição para o backend para reenviar o código.
        // O backend precisaria saber para qual e-mail reenviar o código.
        // Você pode ter passado o e-mail na URL (ex: codigoemail.html?email=...)
        // ou armazenado em sessionStorage da página anterior.

        alert('Solicitando novo código...'); // Apenas para demonstração

        // Exemplo de requisição para o backend (requer um backend):
        /*
        const userEmail = sessionStorage.getItem('resetEmail'); // Exemplo de como pegar o email da página anterior
        if (userEmail) {
            fetch('/api/reenviar-codigo', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email: userEmail })
            })
            .then(response => {
                if (response.ok) {
                    alert('Um novo código foi enviado para seu e-mail!');
                    startCountdown(); // Reinicia o contador após o reenvio
                } else {
                    alert('Não foi possível reenviar o código. Tente novamente.');
                }
            })
            .catch(error => {
                console.error('Erro ao reenviar:', error);
                alert('Erro de conexão ao reenviar código.');
            });
        } else {
            alert('Não foi possível identificar o e-mail para reenvio.');
        }
        */

        // Para a demonstração sem backend, apenas reinicie o contador:
        startCountdown();

        
    });
});
