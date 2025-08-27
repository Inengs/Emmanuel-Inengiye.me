// this is to ensure smooth movement upon clicking of links in the navbar
document.querySelectorAll('nav a').forEach(function (anchor) {
    anchor.addEventListener('click', function (e) {
        e.preventDefault()
        const target = document.querySelector(this.getAttribute('href'));
        target.scrollIntoView({
            behaviour: 'smooth'
        });
    });
});


// const hamburger = document.getElementById("hamburger");
// const navbar = document.getElementById("navbar");

// hamburger.addEventListener("click", () => {
//     navbar.classList.toggle("active");
// });


// Auto-scrolling for projects slider
document.addEventListener('DOMContentLoaded', () => {
    const track = document.querySelector('.track');
    const cards = document.querySelectorAll('.card');
    const cardWidth = cards[0].offsetWidth + 16; // Card width + gap (16px from CSS)
    let intervalId;

    // Function to scroll to the next card
    function scrollNext() {
        const maxScroll = track.scrollWidth - track.clientWidth; // Max scrollable distance
        if (track.scrollLeft >= maxScroll - 5) { // Near the end
            track.scrollTo({ left: 0, behavior: 'smooth' }); // Loop to start
        } else {
            track.scrollBy({ left: cardWidth, behavior: 'smooth' }); // Scroll to next card
        }
    }

    // Start auto-scrolling every 5 seconds
    function startAutoScroll() {
        intervalId = setInterval(scrollNext, 5000); // Scroll every 5 seconds
    }

    // Stop auto-scrolling
    function stopAutoScroll() {
        clearInterval(intervalId);
    }

    // Start auto-scroll when page loads
    startAutoScroll();

    // Pause auto-scroll on hover
    track.addEventListener('mouseenter', stopAutoScroll);
    track.addEventListener('mouseleave', startAutoScroll);

    // Manual controls for Previous and Next buttons
    document.querySelector('.ctrl.next').addEventListener('click', () => {
        stopAutoScroll(); // Pause auto-scroll on manual click
        scrollNext();
        startAutoScroll(); // Resume after click
    });

    document.querySelector('.ctrl.prev').addEventListener('click', () => {
        stopAutoScroll(); // Pause auto-scroll on manual click
        track.scrollBy({ left: -cardWidth, behavior: 'smooth' }); // Scroll to previous card
        startAutoScroll(); // Resume after click
    });
});

document.addEventListener('DOMContentLoaded', () => {
    const faqItems = document.querySelectorAll('.faq-item');

    faqItems.forEach(item => {
        item.querySelector('.faq-question').addEventListener('click', () => {
            item.classList.toggle('active');
        });
    });
})

document.getElementById('contact-form')?.addEventListener('submit', (e) => {
    e.preventDefault();
    const email = e.target.querySelector('input[type="email"]').value;
    alert(`Email submitted: ${email}`); // Replace with actual submission logic
});


const hamburger = document.getElementById("hamburger");
const navbar = document.getElementById("navbar");

if (hamburger && navbar) {
    hamburger.addEventListener("click", () => {
        navbar.classList.toggle("active");
    });
}