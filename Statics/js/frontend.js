document.addEventListener('DOMContentLoaded', () => {
    fetch('/api/paciente')
      .then(res => res.json())
      .then(p => {
        document.getElementById('pacienteNomeCompleto')?.textContent = p.nomeCompleto || '';
        document.getElementById('pacienteDataNascimento')?.textContent = p.dataNascimento || '';
        document.getElementById('pacienteCPF')?.textContent = p.cpf || '';
        document.getElementById('pacienteEmail')?.textContent = p.online?.email || '';
        document.getElementById('pacienteCelular')?.textContent = p.celular || '';
        document.getElementById('pacienteCartaoSUS')?.textContent = p.cartaoSUS || '';
        document.getElementById('pacienteRG')?.textContent = p.rg || '';
        document.getElementById('pacienteSexo')?.textContent = p.sexo || '';
        document.getElementById('pacienteEnderecoCompleto')?.textContent = 
          `${p.endereco?.logradouro || ''}`;
        document.getElementById('pacienteCEP')?.textContent = p.endereco?.cep || '';
        document.getElementById('pacienteTelefone')?.textContent = p.telefone || '';
        document.getElementById('pacienteEscolaridade')?.textContent = p.escolaridade || '';
        document.getElementById('pacienteObservacao')?.textContent = p.observacao || '';
      });

    document.getElementById('btnProntuario')?.addEventListener('click', () => {
        document.getElementById('perfilPacienteContainer').style.display = '';
        document.getElementById('historicoHospitalarContainer').style.display = 'none';
    });

    document.getElementById('btnFichaCitopatologica')?.addEventListener('click', () => {
        document.getElementById('perfilPacienteContainer').style.display = 'none';
        document.getElementById('historicoHospitalarContainer').style.display = '';
    });

    const perfilBtn = document.getElementById('headerIconePerfilTrigger');
    const dropdown = document.getElementById('perfilDropdownOverlay');
    perfilBtn?.addEventListener('click', () => {
        dropdown?.classList.toggle('visivel');
    });

    const cpfEl = document.getElementById('pacienteCPF');
    cpfEl?.addEventListener('input', function () {
        let val = this.value.replace(/\D/g, '').slice(0, 11);
        val = val.replace(/(\d{3})(\d)/, '$1.$2');
        val = val.replace(/(\d{3})(\d)/, '$1.$2');
        val = val.replace(/(\d{3})(\d{1,2})$/, '$1-$2');
        this.value = val;
    });

    const form = document.getElementById('formCadastroPaciente');
    form?.addEventListener('submit', function (e) {
        const nome = document.getElementById('pacienteNome')?.value.trim();
        if (!nome) {
            e.preventDefault();
            alert('Preencha o nome do paciente.');
        }
    });

    const logoutLink = document.getElementById('linkSairDaConta');
    const logoutModal = document.getElementById('confirmLogoutModal');
    const btnConfirmaLogout = document.getElementById('confirmLogoutModalButton');
    logoutLink?.addEventListener('click', () => {
        logoutModal?.classList.add('visivel');
    });
    btnConfirmaLogout?.addEventListener('click', () => {
        localStorage.removeItem('pacienteLogadoIdentificador');
        window.location.href = 'escolher-tipo-usuario.html';
    });
});
