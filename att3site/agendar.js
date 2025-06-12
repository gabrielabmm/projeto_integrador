document.addEventListener('DOMContentLoaded', () => {
    const btnAgendarExames = document.getElementById('btnAgendarExames');
    const dateModal = document.getElementById('dateModal');
    const ubsTimeModal = document.getElementById('ubsTimeModal'); // Get the new modal
const contactModal = document.getElementById('contactModal');
    const detailsModal = document.getElementById('detailsModal');
    const closeContactModalBtn = document.getElementById('closeContactModal');
    const closeDetailsModalBtn = document.getElementById('closeDetailsModal');
    const contactBtns = document.querySelectorAll('.exam-card .contact-btn');
    const detailsBtns = document.querySelectorAll('.exam-card .details-btn');

    const contactModalBody = document.getElementById('contactModalBody');
    const detailsModalBody = document.getElementById('detailsModalBody');
    const currentDayDisplay = document.querySelector('.current-date-display');
    const prevMonthBtn = document.querySelector('.prev-month');
    const nextMonthBtn = document.querySelector('.next-month');
    const monthYearDisplay = document.querySelector('.month-year');
    const calendarGrid = document.querySelector('.calendar-grid');
    const cancelDateSelectionBtn = document.getElementById('cancelDateSelection');
    const okDateSelectionBtn = document.getElementById('okDateSelection');
    const cancelUbsTimeSelectionBtn = document.getElementById('cancelUbsTimeSelection'); // New button
    const confirmAppointmentBtn = document.getElementById('confirmAppointment'); // New button
    const selectUbs = document.getElementById('selectUbs'); // New select element
    const selectTime = document.getElementById('selectTime'); // New select element

    let currentDate = new Date();
    let activeDate = new Date(); // Dia inicialmente selecionado será o dia atual
    let currentMonth = currentDate.getMonth();
    let currentYear = currentDate.getFullYear();

    const monthNames = [
        "Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
        "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"
    ];
    const dayNames = [
        "Domingo", "Segunda-feira", "Terça-feira", "Quarta-feira",
        "Quinta-feira", "Sexta-feira", "Sábado"
    ];

    function openModal(modal) {
        modal.classList.add('active');
    }

    function closeModal(modal) {
        modal.classList.remove('active');
    }

    function renderCalendar() {
        while (calendarGrid.children.length > 7) {
            calendarGrid.removeChild(calendarGrid.lastChild);
        }

        monthYearDisplay.textContent = `${monthNames[currentMonth]} ${currentYear}`;

        const firstDayOfMonth = new Date(currentYear, currentMonth, 1).getDay();
        const lastDayOfMonth = new Date(currentYear, currentMonth + 1, 0).getDate();

        for (let i = 0; i < firstDayOfMonth; i++) {
            const emptyDay = document.createElement('span');
            emptyDay.classList.add('day', 'empty');
            calendarGrid.appendChild(emptyDay);
        }

        for (let day = 1; day <= lastDayOfMonth; day++) {
            const dayElement = document.createElement('span');
            dayElement.classList.add('day');
            dayElement.textContent = day;

            if (day === currentDate.getDate() &&
                currentMonth === currentDate.getMonth() &&
                currentYear === currentDate.getFullYear()) {
                dayElement.classList.add('today');
            }

            if (day === activeDate.getDate() &&
                currentMonth === activeDate.getMonth() &&
                currentYear === activeDate.getFullYear()) {
                dayElement.classList.add('selected');
            }

            dayElement.addEventListener('click', () => {
                const previouslySelected = calendarGrid.querySelector('.day.selected');
                if (previouslySelected) {
                    previouslySelected.classList.remove('selected');
                }
                dayElement.classList.add('selected');

                activeDate = new Date(currentYear, currentMonth, day);
                updateCurrentDateDisplay(activeDate);
            });

            calendarGrid.appendChild(dayElement);
        }
    }

    
    function getFormattedDate(date) {
        const dayOfWeek = dayNames[date.getDay()];
        const dayOfMonth = date.getDate();
        const month = monthNames[date.getMonth()];
        const year = date.getFullYear();
        return `${dayOfWeek}, ${dayOfMonth} de ${month} de ${year}`;
    }

    function updateCurrentDateDisplay(date) {
        const dayOfWeek = dayNames[date.getDay()];
        const dayOfMonth = date.getDate();
        const month = monthNames[date.getMonth()];
        const year = date.getFullYear();
         currentDayDisplay.textContent = getFormattedDate(date);
    }

    if (btnAgendarExames) {
        btnAgendarExames.addEventListener('click', () => {
            renderCalendar();
            updateCurrentDateDisplay(activeDate);
            openModal(dateModal);
        });
    }

    if (prevMonthBtn) {
        prevMonthBtn.addEventListener('click', () => {
            currentMonth--;
            if (currentMonth < 0) {
                currentMonth = 11;
                currentYear--;
            }
            renderCalendar();
        });
    }

    if (nextMonthBtn) {
        nextMonthBtn.addEventListener('click', () => {
            currentMonth++;
            if (currentMonth > 11) {
                currentMonth = 0;
                currentYear++;
            }
            renderCalendar();
        });
    }

    if (cancelDateSelectionBtn) {
        cancelDateSelectionBtn.addEventListener('click', () => {
            closeModal(dateModal);
        });
    }

    if (okDateSelectionBtn) {
        okDateSelectionBtn.addEventListener('click', () => {
            closeModal(dateModal);
            openModal(ubsTimeModal); 
        });
    }

    if (cancelUbsTimeSelectionBtn) {
        cancelUbsTimeSelectionBtn.addEventListener('click', () => {
            closeModal(ubsTimeModal); 
        });
    }

   if (confirmAppointmentBtn) {
        confirmAppointmentBtn.addEventListener('click', () => {
            const selectedUbs = selectUbs.value;
            const selectedTime = selectTime.value;
            const formattedDate = getFormattedDate(activeDate);

            if (selectedUbs && selectedTime) {
                alert(`Agendamento confirmado para:\nData: ${formattedDate}\nUBS: ${selectedUbs}\nHorário: ${selectedTime}`);
                closeModal(ubsTimeModal);
            } else {
                alert('Por favor, selecione a UBS e o horário.');
            }
        });
    }


    if (dateModal) {
        dateModal.addEventListener('click', (e) => {
            if (e.target === dateModal) {
                closeModal(dateModal);
            }
        });
    }

    
    if (ubsTimeModal) {
        ubsTimeModal.addEventListener('click', (e) => {
            if (e.target === ubsTimeModal) {
                closeModal(ubsTimeModal);
            }
        });
    }

    contactBtns.forEach(btn => {
        btn.addEventListener('click', (event) => {
            
            const examCard = event.target.closest('.exam-card');

            
            const laboratorioNome = examCard.dataset.laboratorioNome;
            const laboratorioTelefone = examCard.dataset.laboratorioTelefone;
            const laboratorioEmail = examCard.dataset.laboratorioEmail;
            const laboratorioEndereco = examCard.dataset.laboratorioEndereco;

            contactModalBody.innerHTML = `
                <p><strong>Laboratório:</strong> ${laboratorioNome}</p>
                <p><strong>Telefone:</strong> ${laboratorioTelefone}</p>
                <p><strong>Email:</strong> ${laboratorioEmail}</p>
                <p><strong>Endereço:</strong> ${laboratorioEndereco}</p>
            `;
            openModal(contactModal);
        });
    });

    detailsBtns.forEach(btn => {
        btn.addEventListener('click', (event) => {
            const examCard = event.target.closest('.exam-card');

            const exameTipo = examCard.dataset.exameTipo;
            const medicoSolicitante = examCard.dataset.medicoSolicitante;
            const dataAgendada = examCard.dataset.dataAgendada;
            const horario = examCard.dataset.horario;
            const local = examCard.dataset.local;
            const preparo = examCard.dataset.preparo;
            const observacoes = examCard.dataset.observacoes;

            detailsModalBody.innerHTML = `
                <p><strong>Tipo de Exame:</strong> ${exameTipo}</p>
                <p><strong>Médico Solicitante:</strong> ${medicoSolicitante}</p>
                <p><strong>Data Agendada:</strong> ${dataAgendada}</p>
                <p><strong>Horário:</strong> ${horario}</p>
                <p><strong>Local:</strong> ${local}</p>
                <p><strong>Preparo:</strong> ${preparo}</p>
                <p><strong>Observações:</strong> ${observacoes}</p>
            `;
            openModal(detailsModal);
        });
    });

    if (closeContactModalBtn) {
        closeContactModalBtn.addEventListener('click', () => {
            closeModal(contactModal);
        });
    }

    if (closeDetailsModalBtn) {
        closeDetailsModalBtn.addEventListener('click', () => {
            closeModal(detailsModal);
        });
    }

    if (contactModal) {
        contactModal.addEventListener('click', (e) => {
            if (e.target === contactModal) {
                closeModal(contactModal);
            }
        });
    }

    if (detailsModal) {
        detailsModal.addEventListener('click', (e) => {
            if (e.target === detailsModal) {
                closeModal(detailsModal);
            }
        });
    }

    renderCalendar();
    updateCurrentDateDisplay(activeDate);
});
