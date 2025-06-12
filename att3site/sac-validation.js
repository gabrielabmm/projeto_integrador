document.addEventListener('DOMContentLoaded', function() {
    const sacForm = document.getElementById('sacForm');
    const emailInput = document.getElementById('email');
    const messageTextarea = document.getElementById('mensagem');
    const emailErrorDiv = document.getElementById('emailErrorSac');
    const sacSuccessMessageDiv = document.getElementById('sacSuccessMessage');
    const submitButton = document.getElementById('sacSubmitButton'); // Usando o ID do botão
    const labelEmail = document.getElementById('labelEmail'); // Usando o ID do label
    const labelMensagem = document.getElementById('labelMensagem'); // Usando o ID do label

    if (sacForm && emailInput && messageTextarea && emailErrorDiv && sacSuccessMessageDiv && submitButton && labelEmail && labelMensagem) {
        sacForm.addEventListener('submit', function(event) {
            event.preventDefault(); // Previne o envio padrão do formulário para controlarmos o fluxo

            // Limpa mensagens anteriores
            emailErrorDiv.style.display = 'none';
            emailErrorDiv.textContent = '';
            emailInput.classList.remove('invalid');
            sacSuccessMessageDiv.style.display = 'none'; // Garante que a msg de sucesso esteja oculta inicialmente

            const emailValue = emailInput.value.trim();
            const messageValue = messageTextarea.value.trim(); // Pega o valor da mensagem

            let isValid = true;

            // Validação do email
            if (!isValidEmail(emailValue)) {
                emailErrorDiv.textContent = 'Insira um email válido';
                emailErrorDiv.style.display = 'block';
                emailInput.classList.add('invalid');
                isValid = false;
            }

            // Validação da mensagem (exemplo: não pode estar vazia)
            if (messageValue === '') {
                // Se você quiser um erro específico para a mensagem, adicione um div similar ao emailErrorSac
                // Por enquanto, vamos apenas considerar inválido e focar na mensagem de sucesso.
                // Exemplo: alert('A mensagem não pode estar vazia.');
                console.log('Campo de mensagem vazio'); // Apenas para depuração
                // Você pode adicionar uma mensagem de erro para o campo de mensagem aqui se desejar
                // e então definir isValid = false;
            }

            if (isValid) {
                // SIMULAÇÃO DE ENVIO BEM-SUCEDIDO
                console.log('Formulário SAC válido. Simulando envio...');
                console.log('Email:', emailValue);
                console.log('Mensagem:', messageValue);

                // Oculta os elementos do formulário
                labelEmail.style.display = 'none';
                emailInput.style.display = 'none';
                labelMensagem.style.display = 'none';
                messageTextarea.style.display = 'none';
                submitButton.style.display = 'none';
                emailErrorDiv.style.display = 'none'; // Garante que erro de email também suma

                // Exibe a mensagem de sucesso
                sacSuccessMessageDiv.textContent = 'Mensagem enviada com sucesso!'; // Define o texto aqui, caso queira mudar dinamicamente
                sacSuccessMessageDiv.style.display = 'block';

                // Opcional: Resetar o formulário e mostrar novamente após alguns segundos
                // setTimeout(() => {
                //     sacForm.reset(); // Limpa os campos do formulário
                //     labelEmail.style.display = 'block';
                //     emailInput.style.display = 'block';
                //     labelMensagem.style.display = 'block';
                //     messageTextarea.style.display = 'block';
                //     submitButton.style.display = 'block';
                //     sacSuccessMessageDiv.style.display = 'none';
                // }, 5000); // Reexibe o formulário após 5 segundos

            }
        });
    }

    function isValidEmail(email) {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
    }

    // Limpar erro de email ao digitar
    if (emailInput && emailErrorDiv) {
        emailInput.addEventListener('input', function() {
            if (emailErrorDiv.style.display === 'block') {
                emailErrorDiv.textContent = '';
                emailErrorDiv.style.display = 'none';
                emailInput.classList.remove('invalid');
            }
        });
    }
});
