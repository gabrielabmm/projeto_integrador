document.addEventListener('DOMContentLoaded', () => {
    const btnAgendarExames = document.getElementById('btnAgendarExames');
    const dateModal = document.getElementById('dateModal');
    const ubsTimeModal = document.getElementById('ubsTimeModal');
    const contactModal = document.getElementById('contactModal');
    const detailsModal = document.getElementById('detailsModal');

    const contactModalBody = document.getElementById('contactModalBody');
    const detailsModalBody = document.getElementById('detailsModalBody');
    const currentDayDisplay = document.querySelector('.current-date-display');
    const calendarGrid = document.querySelector('.calendar-grid');

    const cancelDateSelectionBtn = document.getElementById('cancelDateSelection');
    const okDateSelectionBtn = document.getElementById('okDateSelection');
    const cancelUbsTimeSelectionBtn = document.getElementById('cancelUbsTimeSelection');
    const confirmAppointmentBtn = document.getElementById('confirmAppointment');
    const selectUbs = document.getElementById('selectUbs');
    const selectTime = document.getElementById('selectTime');

    let selectedDate = null;

    function openModal(modal) {
        modal.classList.add('active');
    }

    function closeModal(modal) {
        modal.classList.remove('active');
    }

    function loadCalendar() {
        fetch('/api/calendario')
            .then(res => res.json())
            .then(days => {
                while (calendarGrid.children.length > 7) {
                    calendarGrid.removeChild(calendarGrid.lastChild);
                }

                days.forEach(day => {
                    const el = document.createElement('span');
                    el.classList.add('day');
                    if (day.is_empty) {
                        el.classList.add('empty');
                        calendarGrid.appendChild(el);
                        return;
                    }
                    el.textContent = day.day;
                    if (day.is_today) el.classList.add('today');
                    if (day.is_selected) {
                        el.classList.add('selected');
                        selectedDate = day.day;
                        updateDateDisplay(day.day);
                    }

                    el.addEventListener('click', () => {
                        calendarGrid.querySelectorAll('.day.selected').forEach(d => d.classList.remove('selected'));
                        el.classList.add('selected');
                        selectedDate = day.day;
                        updateDateDisplay(day.day);
                    });

                    calendarGrid.appendChild(el);
                });
            });
    }

    function updateDateDisplay(day) {
        const now = new Date();
        now.setDate(day);
        const options = { weekday: 'long', day: 'numeric', month: 'long', year: 'numeric' };
        currentDayDisplay.textContent = now.toLocaleDateString('pt-BR', options);
    }

    if (btnAgendarExames) {
        btnAgendarExames.addEventListener('click', () => {
            openModal(dateModal);
            loadCalendar();
        });
    }

    if (cancelDateSelectionBtn) {
        cancelDateSelectionBtn.addEventListener('click', () => closeModal(dateModal));
    }

    if (okDateSelectionBtn) {
        okDateSelectionBtn.addEventListener('click', () => {
            closeModal(dateModal);
            openModal(ubsTimeModal);
        });
    }

    if (cancelUbsTimeSelectionBtn) {
        cancelUbsTimeSelectionBtn.addEventListener('click', () => closeModal(ubsTimeModal));
    }

    if (confirmAppointmentBtn) {
        confirmAppointmentBtn.addEventListener('click', () => {
            const ubs = selectUbs.value;
            const time = selectTime.value;

            if (!ubs || !time || !selectedDate) {
                alert("Selecione a UBS, o horário e uma data.");
                return;
            }

            const now = new Date();
            now.setDate(selectedDate);
            const formattedDate = now.toLocaleDateString('pt-BR');

            fetch('/api/confirmar', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ date: formattedDate, ubs, time })
            })
                .then(res => res.json())
                .then(data => {
                    console.log(data);
                    alert(`Agendamento confirmado para ${data.data} às ${data.horario} na ${data.ubs}`);
                    closeModal(ubsTimeModal);
                });
        });
    }

    // Exibir dados dos exames
    const renderExamCards = () => {
        fetch('/api/exames')
            .then(res => res.json())
            .then(exames => {
                const container = document.querySelector('.secao-exames-agendados');
                console.log("Exames recebidos:", exames);
                exames.forEach((exame, index) => {
                    const card = document.createElement('div');
                    card.classList.add('exam-card', exame.status === 'cancelled' ? 'cancelled' : 'confirmed');
                    card.dataset.index = index;
                    card.dataset.exame = JSON.stringify(exame);

                    card.innerHTML = `
                        <div class="status-header">
                            <span class="status-dot-${exame.status === 'cancelled' ? 'cancelado' : 'confirmado'}"></span>
                            <p class="status-text">${exame.status === 'cancelled' ? 'Cancelado' : 'Confirmado'}</p>
                            <p class="protocol">Nº de protocolo: ${exame.protocolo}</p>
                        </div>
                        <div class="exam-details">
                            <div class="detail-item">
                                <p class="label">Médico:</p>
                                <p class="value">${exame.medico}</p>
                            </div>
                            <div class="detail-item">
                                <p class="label">Exame:</p>
                                <p class="value">${exame.exame}</p>
                            </div>
                            <div class="detail-item">
                                <p class="label">Laboratório:</p>
                                <p class="value">${exame.laboratorio}</p>
                            </div>
                            <div class="detail-item">
                                <p class="label">Data da coleta:</p>
                                <p class="value">${exame.data}</p>
                            </div>
                        </div>
                        <div class="exam-actions">
                            <button class="btn contact-btn">Contato</button>
                            <button class="btn details-btn">Mais detalhes</button>
                        </div>
                    `;

                    card.querySelector('.contact-btn').addEventListener('click', () => {
                        contactModalBody.innerHTML = `
                            <p><strong>Laboratório:</strong> ${exame.laboratorio}</p>
                            <p><strong>Telefone:</strong> ${exame.telefone}</p>
                            <p><strong>Email:</strong> ${exame.email}</p>
                            <p><strong>Endereço:</strong> ${exame.endereco}</p>
                        `;
                        openModal(contactModal);
                    });

                    card.querySelector('.details-btn').addEventListener('click', () => {
                        console.log("Exame recebido:", exame);
                        detailsModalBody.innerHTML = `
                            <p><strong>Tipo de Exame:</strong> ${exame.exame}</p>
                            <p><strong>Médico Solicitante:</strong> ${exame.medico}</p>
                            <p><strong>Data Agendada:</strong> ${exame.data}</p>
                            <p><strong>Horário:</strong> ${exame.horario}</p>
                            <p><strong>Local:</strong> ${exame.local}</p>
                            <p><strong>Preparo:</strong> ${exame.preparo}</p>
                            <p><strong>Observações:</strong> ${exame.observacoes}</p>
                        `;
                        openModal(detailsModal);
                    });

                    container.appendChild(card);
                });
            });
    };

    renderExamCards();

    document.getElementById('closeContactModal')?.addEventListener('click', () => closeModal(contactModal));
    document.getElementById('closeDetailsModal')?.addEventListener('click', () => closeModal(detailsModal));

    [contactModal, detailsModal, dateModal, ubsTimeModal].forEach(modal => {
        modal?.addEventListener('click', (e) => {
            if (e.target === modal) closeModal(modal);
        });
    });
});
