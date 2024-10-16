window.HELP_IMPROVE_VIDEOJS = false;

$(document).ready(function() {
    // Check for click events on the navbar burger icon
    $(".navbar-burger").click(function() {
      // Toggle the "is-active" class on both the "navbar-burger" and the "navbar-menu"
      $(".navbar-burger").toggleClass("is-active");
      $(".navbar-menu").toggleClass("is-active");

    });

    var options = {
			slidesToScroll: 1,
			slidesToShow: 3,
			loop: true,
			infinite: true,
			autoplay: false,
			pagination: false
    }

		// Initialize all div with carousel class
    var carousels = bulmaCarousel.attach('.carousel', options);

    // Loop on each carousel initialized
    for(var i = 0; i < carousels.length; i++) {
    	// Add listener to  event
    	carousels[i].on('before:show', state => {
    		console.log(state);
    	});
    }

    // Access to bulmaCarousel instance of an element
    var element = document.querySelector('#my-element');
    if (element && element.bulmaCarousel) {
    	// bulmaCarousel instance is available as element.bulmaCarousel
    	element.bulmaCarousel.on('before-show', function(state) {
    		console.log(state);
    	});
    }

    bulmaSlider.attach();

})

document.addEventListener('DOMContentLoaded', function () {
	// Get all tab elements
	var tabs = document.querySelectorAll('.tabs ul li');
	// Get all tab content boxes
	var tabContents = document.querySelectorAll('.tab-content');

	tabs.forEach(function (tab) {
		tab.addEventListener('click', function () {
			var target = this.dataset.tab;

			// Remove 'is-active' class from all tabs
			tabs.forEach(function (tab) {
				tab.classList.remove('is-active');
			});
			// Add 'is-active' class to the clicked tab
			this.classList.add('is-active');

			// Hide all tab contents
			tabContents.forEach(function (content) {
				content.classList.add('is-hidden');
			});
			// Show the content corresponding to the clicked tab
			document.getElementById('tab-content-' + target).classList.remove('is-hidden');
		});
	});
	});
