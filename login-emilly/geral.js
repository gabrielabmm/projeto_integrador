(function() {
    'use strict'; // Ajuda a pegar erros comuns

    // === Conteúdo de cadastro.js (para cadastro.html) ===
    (function() {
        // Condição para executar apenas na página de cadastro
        // Verifica a existência de um elemento chave do formulário de cadastro
        if (document.getElementById('crm') && document.querySelector('form[id="cadastroForm"]')) { // Adicionei um ID ao form para especificidade
            console.log('Executando scripts para: Cadastro');

            const form = document.querySelector('form[id="cadastroForm"]'); // Seja específico se houver múltiplos forms
            const crmInput = document.getElementById('crm');
            const cnesInput = document.getElementById('cnes'); // Se este campo for opcional e puder não existir, ajuste a lógica
            const emailInput = document.getElementById('email'); // Se 'email' for um ID comum, considere um ID mais específico como 'cadastroEmail'
            const senhaInput = document.getElementById('senha'); // Idem para 'senha', considere 'cadastroSenha'
            const sucessoContainer = document.getElementById('sucessoContainer');

            // Certifique-se de que todos os elementos existem antes de adicionar listeners
            if (form && crmInput && emailInput && senhaInput && sucessoContainer) {
                form.addEventListener('submit', function (e) {
                    e.preventDefault();

                    const crm = crmInput.value.trim();
                    const cnes = cnesInput ? cnesInput.value.trim() : ''; // Trata cnesInput opcional
                    const email = emailInput.value.trim();
                    const senha = senhaInput.value.trim();

                    if (!crm || !email || !senha) {
                        alert('Por favor, preencha todos os campos obrigatórios (CRM, Email, Senha).');
                        return;
                    }

                    if (senha.length < 6) {
                        alert('A senha deve ter no mínimo 6 caracteres.');
                        return;
                    }

                    form.style.display = 'none';
                    sucessoContainer.style.display = 'flex';
                });
            } else {
                console.warn('Elementos do formulário de cadastro não encontrados.');
            }

            // Função voltarParaInicio: Se usada por onclick="voltarParaInicio()" no HTML,
            // descomente a linha abaixo. Caso contrário, pode ser interna ou removida se não usada.
            // window.voltarParaInicio = function() {
            //     window.location.href = 'index.html';
            // };
            // Se for um botão específico na página de cadastro com um ID, faça assim:
            // const btnVoltarInicioCadastro = document.getElementById('btnVoltarInicioCadastro');
            // if (btnVoltarInicioCadastro) {
            //     btnVoltarInicioCadastro.addEventListener('click', function() {
            //         window.location.href = 'index.html';
            //     });
            // }

        }
    })();

    // === Conteúdo de codigoemail.js (para codigoemail.html) ===
    (function() {
        // Condição para executar apenas na página de verificação de código
        if (document.getElementById('codigoInput') && document.getElementById('verificarCodigoBtn')) {
            console.log('Executando scripts para: Código Email');

            const codigoInput = document.getElementById('codigoInput');
            const codigoError = document.getElementById('codigoError');
            const codigoSuccessMessage = document.getElementById('codigoSuccessMessage');
            const verificarCodigoBtn = document.getElementById('verificarCodigoBtn');
            const reenviarCodigoBtn = document.getElementById('reenviarCodigoBtn');
            const reenviarInfoSpan = document.getElementById('reenviarInfo'); // Pode ser nulo se o span do contador estiver dentro do botão

            let countdownTime = 60;
            let countdownInterval;

            function isValidCodeFormat(code) {
                const codeRegex = /^\d{6}$/;
                return codeRegex.test(code);
            }

            function showCodigoError(message) {
                if(codigoInput) codigoInput.classList.add('invalid');
                if(codigoError) {
                    codigoError.textContent = message;
                    codigoError.style.display = 'block';
                }
                if(codigoSuccessMessage) codigoSuccessMessage.style.display = 'none';
            }

            function hideCodigoError() {
                if(codigoInput) codigoInput.classList.remove('invalid');
                if(codigoError) {
                    codigoError.textContent = '';
                    codigoError.style.display = 'none';
                }
            }

            function showCodigoSuccess() {
                if(codigoSuccessMessage) codigoSuccessMessage.style.display = 'block';
                hideCodigoError();
            }

            function hideCodigoSuccess() {
                if(codigoSuccessMessage) codigoSuccessMessage.style.display = 'none';
            }

            function startCountdown() {
                if (!reenviarCodigoBtn) return; // Se o botão não existe, não faz nada

                reenviarCodigoBtn.disabled = true;
                reenviarCodigoBtn.innerHTML = `Reenviar Código (<span id="countdown">${countdownTime}</span>s)`;
                if (reenviarInfoSpan) reenviarInfoSpan.textContent = 'Aguarde antes de reenviar.';
                
                if (countdownInterval) {
                    clearInterval(countdownInterval);
                }

                countdownInterval = setInterval(() => {
                    countdownTime--;
                    const countdownSpanElement = reenviarCodigoBtn.querySelector('#countdown');
                    if (countdownSpanElement) countdownSpanElement.textContent = countdownTime;

                    if (countdownTime <= 0) {
                        clearInterval(countdownInterval);
                        reenviarCodigoBtn.disabled = false;
                        reenviarCodigoBtn.textContent = 'Reenviar Código';
                        if (reenviarInfoSpan) reenviarInfoSpan.textContent = '';
                        countdownTime = 60;
                    }
                }, 1000);
            }

            if (reenviarCodigoBtn) { // Só inicia se o botão de reenviar existir
                 startCountdown();
            }

            if (codigoInput) {
                codigoInput.addEventListener('input', function() {
                    const codigoValue = codigoInput.value.trim();
                    if (codigoValue === '') {
                        hideCodigoError();
                        hideCodigoSuccess();
                    } else if (!isValidCodeFormat(codigoValue)) {
                        showCodigoError('O código deve ter 6 dígitos.');
                        hideCodigoSuccess();
                    } else {
                        hideCodigoError();
                    }
                });
            }

            if (verificarCodigoBtn) {
                verificarCodigoBtn.addEventListener('click', function(event) {
                    event.preventDefault();
                    const codigoValue = codigoInput.value.trim();

                    if (codigoValue === '') {
                        showCodigoError('Por favor, insira o código.');
                        hideCodigoSuccess();
                    } else if (!isValidCodeFormat(codigoValue)) {
                        showCodigoError('O código deve ter 6 dígitos.');
                        hideCodigoSuccess();
                    } else {
                        const isCodeCorrectFromServer = (codigoValue === '123456'); // SIMULAÇÃO
                        if (isCodeCorrectFromServer) {
                            hideCodigoError();
                            showCodigoSuccess();
                            // Em vez de 'redefinicao.html', o fluxo comum seria para a página de definir nova senha.
                            // Você mencionou que `nova-senha.js` é para `escolher-nova-senha.html`.
                            // O `redefinir-senha.js` é para `esqueci-senha.html`.
                            // O fluxo geralmente é: esqueci-senha -> codigo-email -> escolher-nova-senha.
                            // Verifique o nome correto da página aqui.
                            window.location.href = 'escolher-nova-senha.html'; // Ajuste conforme o seu fluxo
                        } else {
                            showCodigoError('Código incorreto. Tente novamente.');
                            hideCodigoSuccess();
                        }
                    }
                });
            }

            if (reenviarCodigoBtn) {
                reenviarCodigoBtn.addEventListener('click', function() {
                    alert('Solicitando novo código...'); // Simulação
                    // Lógica de reenvio (ex: fetch para API)
                    // const userEmail = sessionStorage.getItem('resetEmail');
                    // Adicionar lógica real de reenvio aqui se necessário
                    startCountdown(); // Reinicia o contador
                });
            }
        }
    })();

    // === Conteúdo de nova-senha.js (para escolher-nova-senha.html) ===
    (function() {
        // Condição para executar apenas na página de definir nova senha
        if (document.getElementById('novaSenhaForm') && document.getElementById('novaSenhaInput')) {
            console.log('Executando scripts para: Escolher Nova Senha');

            const novaSenhaInput = document.getElementById('novaSenhaInput');
            const confirmaSenhaInput = document.getElementById('confirmaSenhaInput');
            const novaSenhaError = document.getElementById('novaSenhaError');
            const confirmaSenhaError = document.getElementById('confirmaSenhaError');
            const senhaSuccessMessage = document.getElementById('senhaSuccessMessage');
            // const definirSenhaBtn = document.getElementById('definirSenhaBtn'); // Não usado diretamente no código fornecido, o submit é no form
            const novaSenhaForm = document.getElementById('novaSenhaForm');

            function showErrorMessage(element, message) {
                if(element) {
                    element.textContent = message;
                    element.style.display = 'block';
                }
            }

            function hideErrorMessage(element) {
                if(element) {
                    element.textContent = '';
                    element.style.display = 'none';
                }
            }

            function showSuccessMessage(message) {
                if(senhaSuccessMessage) {
                    senhaSuccessMessage.textContent = message;
                    senhaSuccessMessage.style.display = 'block';
                }
                hideErrorMessage(novaSenhaError);
                hideErrorMessage(confirmaSenhaError);
            }

            function hideSuccessMessage() {
                if(senhaSuccessMessage) senhaSuccessMessage.style.display = 'none';
            }
            
            function validatePasswords() {
                const novaSenhaValue = novaSenhaInput.value;
                const confirmaSenhaValue = confirmaSenhaInput.value;

                hideErrorMessage(novaSenhaError);
                hideErrorMessage(confirmaSenhaError);
                novaSenhaInput.classList.remove('invalid');
                confirmaSenhaInput.classList.remove('invalid');
                hideSuccessMessage(); 

                if (novaSenhaValue.length > 0 && novaSenhaValue.length < 6) { 
                    showErrorMessage(novaSenhaError, 'A senha deve ter no mínimo 6 caracteres.');
                    novaSenhaInput.classList.add('invalid');
                }

                if (confirmaSenhaValue.length > 0 && novaSenhaValue !== confirmaSenhaValue) {
                    showErrorMessage(confirmaSenhaError, 'As senhas não coincidem.');
                    confirmaSenhaInput.classList.add('invalid');
                }
            }

            if (novaSenhaInput) novaSenhaInput.addEventListener('input', validatePasswords);
            if (confirmaSenhaInput) confirmaSenhaInput.addEventListener('input', validatePasswords);

            if (novaSenhaForm) {
                novaSenhaForm.addEventListener('submit', function(event) {
                    event.preventDefault(); 

                    const novaSenhaValue = novaSenhaInput.value;
                    const confirmaSenhaValue = confirmaSenhaInput.value;
                    let hasFinalError = false; 

                    hideErrorMessage(novaSenhaError);
                    hideErrorMessage(confirmaSenhaError);
                    novaSenhaInput.classList.remove('invalid');
                    confirmaSenhaInput.classList.remove('invalid');
                    hideSuccessMessage();

                    if (novaSenhaValue === '') {
                        showErrorMessage(novaSenhaError, 'Por favor, insira sua nova senha.');
                        novaSenhaInput.classList.add('invalid');
                        hasFinalError = true;
                    } else if (novaSenhaValue.length < 6) { 
                        showErrorMessage(novaSenhaError, 'A senha deve ter no mínimo 6 caracteres.');
                        novaSenhaInput.classList.add('invalid');
                        hasFinalError = true;
                    }

                    if (confirmaSenhaValue === '') {
                        showErrorMessage(confirmaSenhaError, 'Por favor, confirme sua senha.');
                        confirmaSenhaInput.classList.add('invalid');
                        hasFinalError = true;
                    } else if (novaSenhaValue !== confirmaSenhaValue) {
                        showErrorMessage(confirmaSenhaError, 'As senhas não coincidem.');
                        confirmaSenhaInput.classList.add('invalid');
                        hasFinalError = true;
                    }

                    if (hasFinalError) {
                        return; 
                    }

                    showSuccessMessage('Senha redefinida com sucesso!');
                    alert('Senha redefinida com sucesso! Redirecionando...');
                    
                    // O script original redirecionava para 'nova-senha.html'.
                    // Se 'escolher-nova-senha.html' é onde o form está, para onde deve ir após o sucesso?
                    // Talvez para uma página de "sucesso" final, ou para a página de login.
                    // Ajuste 'pagina-de-sucesso-final.html' para o destino correto.
                    setTimeout(() => {
                        window.location.href = 'redefinida-sucesso.html'; // Ou login.html, ou uma página de sucesso dedicada
                    }, 1500);
                });
            }
        }
    })();

    // === Conteúdo de redefinir-senha.js (para esqueci-senha.html) ===
    (function() {
        // Condição para executar apenas na página de "esqueci a senha" (onde se insere o email)
        if (document.getElementById('emailForm') && document.getElementById('emailInput')) {
            console.log('Executando scripts para: Esqueci Senha (Redefinir Senha)');

            const emailInput = document.getElementById('emailInput'); 
            const emailError = document.getElementById('emailError');
            const emailSuccessMessage = document.getElementById('emailSuccessMessage'); 
            // const enviarCodigoBtn = document.getElementById('enviarCodigoBtn'); // O listener é no form
            const emailForm = document.getElementById('emailForm'); 

            function isValidEmail(email) {
                const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
                return emailRegex.test(email);
            }

            function showEmailError(message) {
                if(emailInput) emailInput.classList.add('invalid');
                if(emailError) {
                    emailError.textContent = message;
                    emailError.style.display = 'block';
                }
                if(emailSuccessMessage) emailSuccessMessage.style.display = 'none'; 
            }

            function hideEmailError() {
                if(emailInput) emailInput.classList.remove('invalid');
                if(emailError) {
                    emailError.textContent = '';
                    emailError.style.display = 'none';
                }
            }

            function showEmailSuccess(message) {
                if(emailSuccessMessage) {
                    emailSuccessMessage.textContent = message;
                    emailSuccessMessage.style.display = 'block';
                }
                hideEmailError();
            }

            if (emailInput) {
                emailInput.addEventListener('input', function() {
                    const emailValue = emailInput.value.trim();
                    if (emailValue === '') {
                        hideEmailError();
                        if(emailSuccessMessage) emailSuccessMessage.style.display = 'none'; 
                    } else if (!isValidEmail(emailValue)) {
                        showEmailError('Por favor, insira um e-mail válido.');
                    } else {
                        hideEmailError(); 
                        if(emailSuccessMessage) emailSuccessMessage.style.display = 'none'; 
                    }
                });
            }

            if (emailForm) {
                emailForm.addEventListener('submit', function(event) {
                    event.preventDefault(); 
                    const emailValue = emailInput.value.trim();

                    if (emailValue === '') {
                        showEmailError('Por favor, insira seu e-mail.');
                        return; 
                    }
                    if (!isValidEmail(emailValue)) {
                        showEmailError('Por favor, insira um e-mail válido.');
                        return; 
                    }

                    showEmailSuccess('E-mail válido! Enviando código...'); 
                    alert(`Um código de verificação foi enviado para ${emailValue}.`); 
                    sessionStorage.setItem('resetEmail', emailValue); 
                    window.location.href = 'codigoemail.html'; 
                });
            }
        }
    })();

    // === Conteúdo de script.js (POTENCIALMENTE para escolher-tipo-usuario.html ou outra página com form de email) ===
    (function() {
        // Condição para executar: verifica se os elementos específicos deste script existem
        const emailInputGlobal = document.getElementById('email'); // ID diferente do redefinir-senha.js
        const enviarCodigoBtnGlobal = document.getElementById('enviarCodigoBtn'); // Mesmo ID de botão do redefinir-senha.js

        if (emailInputGlobal && enviarCodigoBtnGlobal) {
            console.log('Executando scripts para: Script.js (escolher-tipo-usuario? ou login?)');

            const emailErrorGlobal = document.getElementById('emailError'); // Precisa existir no HTML desta página

            function isValidEmailGlobal(email) {
                const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
                return emailRegex.test(email);
            }

            function showEmailErrorGlobal(message) {
                if(emailInputGlobal) emailInputGlobal.classList.add('invalid');
                if(emailErrorGlobal) {
                    emailErrorGlobal.textContent = message;
                    emailErrorGlobal.style.display = 'block';
                }
            }

            function hideEmailErrorGlobal() {
                if(emailInputGlobal) emailInputGlobal.classList.remove('invalid');
                if(emailErrorGlobal) {
                    emailErrorGlobal.textContent = '';
                    emailErrorGlobal.style.display = 'none';
                }
            }
            
            // A função hideEmailSuccess() era chamada no script original mas não definida.
            // Se precisar dela, defina-a aqui ou certifique-se de que é uma função global.
            // function hideEmailSuccessGlobal() { /* ... lógica ... */ }


            emailInputGlobal.addEventListener('input', function() {
                const emailValue = emailInputGlobal.value.trim();
                if (emailValue === '') {
                    hideEmailErrorGlobal();
                } else if (!isValidEmailGlobal(emailValue)) {
                    showEmailErrorGlobal('Por favor, insira um e-mail válido (ex: seu.email@exemplo.com).');
                } else {
                    hideEmailErrorGlobal();
                }
            });

            enviarCodigoBtnGlobal.addEventListener('click', function(event) {
                event.preventDefault();
                const emailValue = emailInputGlobal.value.trim();

                if (emailValue === '') {
                    showEmailErrorGlobal('O campo de e-mail é obrigatório.');
                } else if (!isValidEmailGlobal(emailValue)) {
                    showEmailErrorGlobal('Por favor, insira um e-mail válido (ex: seu.email@exemplo.com).');
                } else {
                    hideEmailErrorGlobal();
                    // hideEmailSuccess(); // Esta função não está definida aqui. Removi a chamada.
                                        // Se houver uma mensagem de sucesso para esconder, adicione a lógica ou defina a função.
                    console.log('E-mail válido em script.js, redirecionando para codigoemail.html');
                    window.location.href = 'codigoemail.html';
                }
            });
        } else {
            // Este console.log ajudará a identificar se o script.js não encontrou seus elementos
            // console.log('Elementos para script.js (ex: #email, #enviarCodigoBtn) não encontrados na página atual.');
        }
    })();

})(); // Fecha a IIFE principal
