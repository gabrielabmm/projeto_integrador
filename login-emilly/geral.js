(function() {
    'use strict';

    const dadosPacienteAutenticado = {
        nomeCompleto: "Maria Linda da Silva",
        primeiroNome: "Maria",
        dataNascimento: "01/01/1985",
        endereco: {
            cep: "12345-678",
            cidade: "Exemplópolis",
            logradouro: "Rua dos bobos, n°0"
        },
        cpf: "012.345.678-90",
        cartaoSUS: "123.4567.7890.1234",
        online: {
            email: "maria88@linda.com.br",
            login: "linda.silva",
            senha: "Pi2025"
        },
        rg: "00.000.000-0",
        sexo: "Feminino",
        celular: "(11) 98765-4321",
        telefone: "(11) 1234-5678",
        escolaridade: "Superior Completo"
    };

    const dadosInstituicaoAutenticada = {
    nomeInstituicao: "Dra. Joana",
    cnes: "1234567",
    crmCorenResponsavel: "CRM/SP 123456",
    cnpj: "12.345.678/0001-99",
    email: "contato@clinicajoana.com",
    celular: "(11) 91234-5678",
    telefone: "(11) 5555-1234",
    cep: "01000-001",
    cidade: "São Paulo",
    endereco: "Rua da Clínica, Bairro Saúde",
    numero: "100",
    complemento: "Andar 5, Sala 502",
    bairro: "Saúde",
    online: {
        senha: "Pi2025"
    }
};

    function calcularIdade(dataNascimentoString) {
        if (!dataNascimentoString || !/^\d{2}\/\d{2}\/\d{4}$/.test(dataNascimentoString)) { return "N/D"; }
        const partesData = dataNascimentoString.split('/');
        const diaNasc = parseInt(partesData[0], 10);
        const mesNasc = parseInt(partesData[1], 10) - 1;
        const anoNasc = parseInt(partesData[2], 10);
        const dataNascDate = new Date(anoNasc, mesNasc, diaNasc);
        if (isNaN(dataNascDate.getTime())) return "N/D";
        const hoje = new Date();
        let idade = hoje.getFullYear() - dataNascDate.getFullYear();
        const m = hoje.getMonth() - dataNascDate.getMonth();
        if (m < 0 || (m === 0 && hoje.getDate() < dataNascDate.getDate())) { idade--; }
        return idade >= 0 ? idade : "N/D";
    }

    const headerIconePerfilTrigger = document.getElementById('headerIconePerfilTrigger');
    const perfilDropdownOverlay = document.getElementById('perfilDropdownOverlay');
    const fecharDropdownBtn = document.getElementById('fecharDropdownBtn');
    const dropdownFotoPerfil = document.getElementById('dropdownFotoPerfil');
    const dropdownNomeCompleto = document.getElementById('dropdownNomeCompleto');
    const dropdownDataNascimento = document.getElementById('dropdownDataNascimento');
    const dropdownIdade = document.getElementById('dropdownIdade');
    const dropdownCEP = document.getElementById('dropdownCEP');
    const dropdownCPF = document.getElementById('dropdownCPF');
    const dropdownCartaoSUS = document.getElementById('dropdownCartaoSUS');

    const linkSairDaConta = document.getElementById('linkSairDaConta');
    const confirmLogoutModalElement = document.getElementById('confirmLogoutModal');
    const confirmLogoutModalButtonElement = document.getElementById('confirmLogoutModalButton');
    const cancelLogoutModalButtonElement = document.getElementById('cancelLogoutModalButton');

    function togglePerfilDropdown() {
        if (!perfilDropdownOverlay || !dadosPacienteAutenticado) return;
        const isVisivel = perfilDropdownOverlay.classList.contains('visivel');
        if (isVisivel) {
            perfilDropdownOverlay.classList.remove('visivel');
        } else {
            if (dropdownNomeCompleto) dropdownNomeCompleto.textContent = dadosPacienteAutenticado.nomeCompleto || "Nome não disponível";
            if (dropdownDataNascimento) dropdownDataNascimento.textContent = dadosPacienteAutenticado.dataNascimento || "N/D";
            if (dropdownIdade) {
                const idade = calcularIdade(dadosPacienteAutenticado.dataNascimento);
                dropdownIdade.textContent = idade + (idade !== "N/D" ? " anos" : "");
            }
            if (dropdownCEP && dadosPacienteAutenticado.endereco) dropdownCEP.textContent = dadosPacienteAutenticado.endereco.cep || "N/D";
            else if (dropdownCEP) dropdownCEP.textContent = "N/D";
            if (dropdownCPF) dropdownCPF.textContent = dadosPacienteAutenticado.cpf || "N/D";
            if (dropdownCartaoSUS) dropdownCartaoSUS.textContent = dadosPacienteAutenticado.cartaoSUS || "N/D";
            
            perfilDropdownOverlay.classList.add('visivel');
        }
    }

    if (headerIconePerfilTrigger) {
        headerIconePerfilTrigger.addEventListener('click', function(event) {
            event.preventDefault(); 
            event.stopPropagation();
            togglePerfilDropdown();
        });
    }

    if (fecharDropdownBtn) {
        fecharDropdownBtn.addEventListener('click', togglePerfilDropdown);
    }

    document.addEventListener('click', function(event) {
        if (perfilDropdownOverlay && perfilDropdownOverlay.classList.contains('visivel')) {
            if (headerIconePerfilTrigger && !perfilDropdownOverlay.contains(event.target) && !headerIconePerfilTrigger.contains(event.target)) {
                perfilDropdownOverlay.classList.remove('visivel');
            }
        }
    });
    
    document.addEventListener('keydown', function(event) {
        if (event.key === 'Escape' && perfilDropdownOverlay && perfilDropdownOverlay.classList.contains('visivel')) {
            perfilDropdownOverlay.classList.remove('visivel');
        }
    });

    if (linkSairDaConta && confirmLogoutModalElement && confirmLogoutModalButtonElement && cancelLogoutModalButtonElement) {
        linkSairDaConta.addEventListener('click', function(event) {
            event.preventDefault();
            if (perfilDropdownOverlay && perfilDropdownOverlay.classList.contains('visivel')) {
                perfilDropdownOverlay.classList.remove('visivel');
            }
            confirmLogoutModalElement.classList.add('visivel');
        });

        cancelLogoutModalButtonElement.addEventListener('click', function() {
            confirmLogoutModalElement.classList.remove('visivel');
        });

        confirmLogoutModalButtonElement.addEventListener('click', function() {
            localStorage.removeItem('pacienteLogadoIdentificador');
            window.location.href = "escolher-tipo-usuario.html";
        });

        confirmLogoutModalElement.addEventListener('click', function(event) {
            if (event.target === confirmLogoutModalElement) {
                confirmLogoutModalElement.classList.remove('visivel');
            }
        });

        document.addEventListener('keydown', function(event) {
            if (event.key === 'Escape' && confirmLogoutModalElement.classList.contains('visivel')) {
                confirmLogoutModalElement.classList.remove('visivel');
            }
        });
    }

    function preencherFormularioInterna() {
        const dados = dadosPacienteAutenticado;
        const tituloNomeElemento = document.getElementById('titulo-nome');
        if (tituloNomeElemento) {
            tituloNomeElemento.textContent = dados.primeiroNome || (dados.nomeCompleto ? dados.nomeCompleto.split(' ')[0] : 'Usuário');
        }
        const nomePerfilElemento = document.getElementById('nome-perfil');
        if (nomePerfilElemento) {
            nomePerfilElemento.textContent = dados.nomeCompleto || 'Nome não disponível';
        }
        const nomeInput = document.getElementById('nome');
        const emailInput = document.getElementById('email');
        const cpfInput = document.getElementById('cpf');
        const celularInput = document.getElementById('celular');
        const telefoneInput = document.getElementById('telefone');
        const escolaridadeInput = document.getElementById('escolaridade');
        const cepInput = document.getElementById('cep');
        const cidadeInput = document.getElementById('cidade');
        const enderecoInput = document.getElementById('endereco');

        if (nomeInput) nomeInput.value = dados.nomeCompleto || '';
        if (emailInput && dados.online) emailInput.value = dados.online.email || '';
        if (cpfInput) cpfInput.value = dados.cpf || '';
        if (celularInput) celularInput.value = dados.celular || '';
        if (telefoneInput) telefoneInput.value = dados.telefone || '';
        if (escolaridadeInput) escolaridadeInput.value = dados.escolaridade || '';
        if (cepInput && dados.endereco) cepInput.value = dados.endereco.cep || '';
        if (cidadeInput && dados.endereco) cidadeInput.value = dados.endereco.cidade || '';
        if (enderecoInput && dados.endereco) enderecoInput.value = dados.endereco.logradouro || '';

        const camposDoFormulario = document.querySelectorAll('.info-formulario input');
        const botaoEditar = document.querySelector('.info-formulario .botao.editar');
        const infoFormularioDiv = document.querySelector('.info-formulario');

        if (!infoFormularioDiv) return;

        let botoesAcaoContainer = infoFormularioDiv.querySelector('.acoes-formulario');
        if (botoesAcaoContainer) botoesAcaoContainer.remove();

        botoesAcaoContainer = document.createElement('div');
        botoesAcaoContainer.className = 'acoes-formulario';
        botoesAcaoContainer.style.marginTop = '15px';
        botoesAcaoContainer.style.display = 'flex';
        botoesAcaoContainer.style.gap = '10px';
        botoesAcaoContainer.innerHTML = `
            <button type="submit" class="botao salvar">Salvar Alterações</button>
            <button type="button" class="botao cancelar-edicao">Cancelar</button>
        `;
        if (botaoEditar && botaoEditar.parentNode === infoFormularioDiv) {
             infoFormularioDiv.insertBefore(botoesAcaoContainer, botaoEditar);
        } else {
             infoFormularioDiv.appendChild(botoesAcaoContainer);
        }

        const botaoSalvar = botoesAcaoContainer.querySelector('.salvar');
        const botaoCancelarEdicao = botoesAcaoContainer.querySelector('.cancelar-edicao');

        camposDoFormulario.forEach(input => input.disabled = true);
        if(botaoEditar) botaoEditar.style.display = 'inline-block';
        botoesAcaoContainer.style.display = 'none';

        if (botaoEditar) {
            botaoEditar.onclick = function() {
                camposDoFormulario.forEach(input => input.disabled = false);
                botaoEditar.style.display = 'none';
                botoesAcaoContainer.style.display = 'flex';
                if (nomeInput) nomeInput.focus();
            };
        } else {
             camposDoFormulario.forEach(input => input.disabled = false);
             botoesAcaoContainer.style.display = 'flex';
        }

        if (botaoSalvar) {
            botaoSalvar.onclick = function(event) {
                event.preventDefault();
                if(nomeInput) dadosPacienteAutenticado.nomeCompleto = nomeInput.value;
                if(emailInput && dadosPacienteAutenticado.online) dadosPacienteAutenticado.online.email = emailInput.value;
                if(cpfInput) dadosPacienteAutenticado.cpf = cpfInput.value;
                if(celularInput) dadosPacienteAutenticado.celular = celularInput.value;
                if(telefoneInput) dadosPacienteAutenticado.telefone = telefoneInput.value;
                if(escolaridadeInput) dadosPacienteAutenticado.escolaridade = escolaridadeInput.value;
                if(dadosPacienteAutenticado.endereco) {
                    if(cepInput) dadosPacienteAutenticado.endereco.cep = cepInput.value;
                    if(cidadeInput) dadosPacienteAutenticado.endereco.cidade = cidadeInput.value;
                    if(enderecoInput) dadosPacienteAutenticado.endereco.logradouro = enderecoInput.value;
                }
                alert('Dados salvos!');
                window.location.href = "inicio.html";
            };
        }

        if (botaoCancelarEdicao) {
            botaoCancelarEdicao.onclick = function() {
                alert('Edição cancelada.');
                if (typeof window.preencherFormulario === 'function') {
                    window.preencherFormulario();
                }
            };
        }

        const fotoPerfilElement = document.querySelector('.cartao .perfil .avatar img');
        if (fotoPerfilElement) {
            fotoPerfilElement.style.cursor = 'pointer';
            fotoPerfilElement.title = 'Clique para alterar a foto';
            fotoPerfilElement.addEventListener('click', function() {
                const desejaAlterar = window.confirm("Deseja alterar a foto de perfil?");
                if (desejaAlterar) {
                    alert("Funcionalidade para alterar a foto acionada.\n(Em uma implementação real, aqui abriria um seletor de arquivos.)");
                }
            });
        }
    }
    window.preencherFormulario = preencherFormularioInterna;

    function preencherFormularioInstituicaoInterna() {
    const nomeInst = dadosInstituicaoAutenticada.nomeInstituicao || "Nome da Instituição";

    const saudacaoNomeEl = document.getElementById('saudacao-nome');
    if (saudacaoNomeEl) {
        saudacaoNomeEl.textContent = nomeInst;
    }

    const nomeExibicaoPerfilEl = document.getElementById('nome-exibicao-perfil');
    if (nomeExibicaoPerfilEl) {
        nomeExibicaoPerfilEl.textContent = nomeInst;
    }

    const idsDosCamposParaPreencher = [
        'nomeInstituicao', 'cnes', 'crmCorenResponsavel', 'cnpj',
        'email', 'celular', 'telefone', 'cep', 'cidade',
        'endereco', 'numero', 'complemento', 'bairro'
    ];

    idsDosCamposParaPreencher.forEach(id => {
        const el = document.getElementById(id);
        if (el) {
            el.value = dadosInstituicaoAutenticada[id] || '';
        }
    });

    const camposDoFormulario = document.querySelectorAll('.info-formulario input[type="text"], .info-formulario input[type="email"], .info-formulario input[type="tel"]');
    const botaoEditar = document.querySelector('.info-formulario .botao.editar');
    const botaoSalvar = document.querySelector('.info-formulario .botao.salvar');

    if (botaoEditar && botaoSalvar && camposDoFormulario.length > 0) {
        camposDoFormulario.forEach(input => input.disabled = true);
        botaoEditar.style.display = 'inline-block';
        botaoSalvar.style.display = 'none';

        botaoEditar.onclick = function() {
            camposDoFormulario.forEach(input => input.disabled = false);
            botaoEditar.style.display = 'none';
            botaoSalvar.style.display = 'inline-block';
            camposDoFormulario[0].focus();
        };

        botaoSalvar.onclick = function(event) {
            event.preventDefault();

            idsDosCamposParaPreencher.forEach(id => {
                const inputElement = document.getElementById(id);
                if (inputElement && dadosInstituicaoAutenticada.hasOwnProperty(id)) {
                    dadosInstituicaoAutenticada[id] = inputElement.value;
                } else if (inputElement) {
                    dadosInstituicaoAutenticada[id] = inputElement.value;
                }
            });

            alert('Dados da instituição salvos!');
              window.location.href = "iniciomedico.html";
        
            camposDoFormulario.forEach(input => input.disabled = true);
            botaoSalvar.style.display = 'none';
            botaoEditar.style.display = 'inline-block';
        };
    } else {
        if (!botaoEditar) console.error("Botão Editar não encontrado para o formulário da instituição.");
        if (!botaoSalvar) console.error("Botão Salvar não encontrado para o formulário da instituição.");
        if (camposDoFormulario.length === 0) console.error("Nenhum campo de formulário encontrado para a instituição.");
    }

    const fotoPerfilElement = document.querySelector('.cartao .perfil .avatar img');
    if (fotoPerfilElement) {
        fotoPerfilElement.style.cursor = 'pointer';
        fotoPerfilElement.title = 'Clique para alterar a foto';
        fotoPerfilElement.addEventListener('click', function() {
            const desejaAlterar = window.confirm("Deseja alterar a foto de perfil da instituição/profissional?");
            if (desejaAlterar) {
                alert("Funcionalidade para alterar a foto acionada.\n(Em uma implementação real, aqui abriria um seletor de arquivos.)");
            }
        });
    }
}
window.preencherFormularioInstituicao = preencherFormularioInstituicaoInterna;

    (function() {
        if (document.getElementById('crm') && document.querySelector('form[id="cadastroForm"]')) {
        }
    })();

    (function() {
        const codigoInput = document.getElementById('codigoInput');
        const codigoError = document.getElementById('codigoError');
        const codigoSuccessMessage = document.getElementById('codigoSuccessMessage');
        const reenviarCodigoBtn = document.getElementById('reenviarCodigoBtn');
        const countdownSpan = document.getElementById('countdown');
        const formVerificarCodigo = document.getElementById('form-verificar-codigo');

        if (formVerificarCodigo && codigoInput && reenviarCodigoBtn && countdownSpan) {
            let countdownInterval;
            let tempoRestante = 60;

            function iniciarContadorReenvio() {
                tempoRestante = 60;
                reenviarCodigoBtn.disabled = true;
                countdownSpan.textContent = tempoRestante;
                if (countdownInterval) clearInterval(countdownInterval);
                countdownInterval = setInterval(() => {
                    tempoRestante--;
                    countdownSpan.textContent = tempoRestante;
                    if (tempoRestante <= 0) {
                        clearInterval(countdownInterval);
                        reenviarCodigoBtn.disabled = false;
                        countdownSpan.textContent = '';
                    }
                }, 1000);
            }
            iniciarContadorReenvio();

            reenviarCodigoBtn.addEventListener('click', function() {
                alert('Código reenviado! (Simulação)');
                iniciarContadorReenvio();
            });

            formVerificarCodigo.addEventListener('submit', function(event) {
                event.preventDefault();
                if(codigoError) codigoError.textContent = '';
                if(codigoSuccessMessage) codigoSuccessMessage.style.display = 'none';

                const codigoDigitado = codigoInput.value.trim();
                const codigoCorreto = "123456";

                if (codigoDigitado === "") {
                    if(codigoError) codigoError.textContent = "Por favor, insira o código.";
                    if(codigoError && codigoError.style) codigoError.style.display = 'block';
                    codigoInput.focus();
                    return;
                }

                if (codigoDigitado === codigoCorreto) {
                    if(codigoSuccessMessage) {
                        codigoSuccessMessage.textContent = 'Código válido! Redirecionando...';
                        codigoSuccessMessage.style.display = 'block';
                    }
                    if(codigoError && codigoError.style) codigoError.style.display = 'none';
                    setTimeout(function() {
                        window.location.href = 'escolher-nova-senha.html';
                    }, 1500);
                } else {
                    if(codigoError) codigoError.textContent = "Código incorreto. Tente novamente.";
                    if(codigoError && codigoError.style) codigoError.style.display = 'block';
                    codigoInput.focus();
                }
            });
        }
    })();

    (function() {
        const formElement = document.getElementById('novaSenhaForm');
        const novaSenhaInputElement = document.getElementById('novaSenhaInput');

        if (formElement && novaSenhaInputElement) {
            const confirmaSenhaInputElement = document.getElementById('confirmaSenhaInput');
            const novaSenhaErrorElement = document.getElementById('novaSenhaError');
            const confirmaSenhaErrorElement = document.getElementById('confirmaSenhaError');

            formElement.addEventListener('submit', function(event) {
                event.preventDefault();
                if (novaSenhaErrorElement) {
                    novaSenhaErrorElement.textContent = '';
                    novaSenhaErrorElement.style.display = 'none';
                }
                if (confirmaSenhaErrorElement) {
                    confirmaSenhaErrorElement.textContent = '';
                    confirmaSenhaErrorElement.style.display = 'none';
                }

                const novaSenhaValue = novaSenhaInputElement.value;
                const confirmaSenhaValue = confirmaSenhaInputElement ? confirmaSenhaInputElement.value : ''; 
                
                let isFormValid = true;
                const MIN_SENHA_LENGTH = 6;

                if (novaSenhaValue === '') {
                    if (novaSenhaErrorElement) {
                        novaSenhaErrorElement.textContent = 'Por favor, insira a nova senha.';
                        novaSenhaErrorElement.style.display = 'block';
                    } else {
                        alert("Por favor, insira a nova senha.");
                    }
                    isFormValid = false;
                } else if (novaSenhaValue.length < MIN_SENHA_LENGTH) {
                    if (novaSenhaErrorElement) {
                        novaSenhaErrorElement.textContent = `A nova senha deve ter no mínimo ${MIN_SENHA_LENGTH} caracteres.`;
                        novaSenhaErrorElement.style.display = 'block';
                    } else {
                        alert(`A nova senha deve ter no mínimo ${MIN_SENHA_LENGTH} caracteres.`);
                    }
                    isFormValid = false;
                }

                if (novaSenhaValue !== '' && confirmaSenhaValue === '') {
                    if (confirmaSenhaErrorElement) {
                        confirmaSenhaErrorElement.textContent = 'Por favor, confirme a nova senha.';
                        confirmaSenhaErrorElement.style.display = 'block';
                    } else {
                        alert("Por favor, confirme a nova senha.");
                    }
                    isFormValid = false; 
                }

                if (novaSenhaValue !== '' && novaSenhaValue.length >= MIN_SENHA_LENGTH && confirmaSenhaValue !== '') {
                    if (novaSenhaValue !== confirmaSenhaValue) {
                        if (confirmaSenhaErrorElement) {
                            confirmaSenhaErrorElement.textContent = 'As senhas não coincidem.';
                            confirmaSenhaErrorElement.style.display = 'block';
                        } else {
                            alert("As senhas não coincidem.");
                        }
                        isFormValid = false;
                    }
                }
                
                if (isFormValid) {
                    if (typeof dadosPacienteAutenticado !== 'undefined' && dadosPacienteAutenticado && dadosPacienteAutenticado.online) {
                        dadosPacienteAutenticado.online.senha = novaSenhaValue;
                    }
                    window.location.href = 'redefinida-sucesso.html';
                } else {
                    if (novaSenhaErrorElement && novaSenhaErrorElement.style.display === 'block') {
                        novaSenhaInputElement.focus();
                    } else if (confirmaSenhaErrorElement && confirmaSenhaErrorElement.style.display === 'block' && confirmaSenhaInputElement) {
                        confirmaSenhaInputElement.focus();
                    } else {
                        novaSenhaInputElement.focus();
                    }
                }
            });
        }
    })();

    (function() {
        if (document.getElementById('emailForm') && document.getElementById('emailInput')) {
        }
    })();

    (function() {
    })();

    (function() {
        const formLoginUsuario = document.getElementById('formLoginUsuario');
        if (formLoginUsuario) {
            const mensagemErroLoginEl = document.getElementById('mensagem-login-erro');
            formLoginUsuario.addEventListener('submit', function(event) {
                event.preventDefault();
                const tipoSelecionadoRadio = document.querySelector('input[name="tipo"]:checked');
                if (!tipoSelecionadoRadio) {
                    if (mensagemErroLoginEl) {
                        mensagemErroLoginEl.textContent = "Por favor, selecione o tipo de usuário.";
                        mensagemErroErroLoginEl.style.display = 'block';
                    }
                    return;
                }
                const tipoSelecionado = tipoSelecionadoRadio.value;
                if (mensagemErroLoginEl) {
                    mensagemErroLoginEl.textContent = '';
                    mensagemErroLoginEl.style.display = 'none';
                }
                if (tipoSelecionado === 'beneficiario') {
                    const loginInput = document.getElementById('beneficiario-login');
                    const senhaInput = document.getElementById('beneficiario-senha');
                    if (!loginInput || !senhaInput) { console.error("Inputs beneficiário não encontrados"); return; }
                    const loginDigitado = loginInput.value.trim();
                    const senhaDigitada = senhaInput.value.trim();
                    if (loginDigitado === "" || senhaDigitada === "") {
                        if(mensagemErroLoginEl) {
                           mensagemErroLoginEl.textContent = "Por favor, preencha login e senha.";
                           mensagemErroLoginEl.style.display = 'block';
                        }
                        return;
                    }
                    if (dadosPacienteAutenticado && dadosPacienteAutenticado.online &&
                        loginDigitado === dadosPacienteAutenticado.online.login &&
                        senhaDigitada === dadosPacienteAutenticado.online.senha) {
                        localStorage.setItem('pacienteLogadoIdentificador', dadosPacienteAutenticado.online.login);
                        window.location.href = "INICIO.html";
                    } else {
                        if(mensagemErroLoginEl) {
                           mensagemErroLoginEl.textContent = "Login ou senha incorretos para beneficiário.";
                           mensagemErroLoginEl.style.display = 'block';
                        }
                    }
                } else if (tipoSelecionado === 'instituicao') {
                    const cnesInput = document.getElementById('instituicao-cnes');
                    const crmInput = document.getElementById('instituicao-crm');
                    const senhaInstInput = document.getElementById('instituicao-senha-inst');
                    if (!cnesInput || !crmInput || !senhaInstInput) { console.error("Inputs instituição não encontrados"); return; }
                    const cnesDigitado = cnesInput.value.trim();
                    const crmDigitado = crmInput.value.trim();
                    const senhaInstDigitada = senhaInstInput.value.trim();
                    if (cnesDigitado === "" || crmDigitado === "" || senhaInstDigitada === "") {
                           if(mensagemErroLoginEl) {
                               mensagemErroLoginEl.textContent = "Por favor, preencha CNES, CRM/COREN e Senha.";
                               mensagemErroLoginEl.style.display = 'block';
                           }
                           return;
                    }
                    if (dadosInstituicaoAutenticada && dadosInstituicaoAutenticada.online &&
                        cnesDigitado === dadosInstituicaoAutenticada.cnes &&
                        crmDigitado === dadosInstituicaoAutenticada.crmCorenResponsavel &&
                        senhaInstDigitada === dadosInstituicaoAutenticada.online.senha) {
                        localStorage.setItem('instituicaoLogadaIdentificador', dadosInstituicaoAutenticada.cnes);
                        window.location.href = "iniciomedico.html";
                    } else {
                        if(mensagemErroLoginEl) {
                           mensagemErroLoginEl.textContent = "CNES, CRM/COREN ou Senha incorretos para instituição.";
                           mensagemErroLoginEl.style.display = 'block';
                        }
                    }
                }
            });
        }
    })();

})();
