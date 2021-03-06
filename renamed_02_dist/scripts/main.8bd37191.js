// Generated by CoffeeScript 1.6.2
/*!
#
#  Web Starter Kit
#  Copyright 2014 Google Inc. All rights reserved.
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#    https://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License
#
*/


(function() {
  (function() {
    var appbarElement, body, closeMenu, main, menuBtn, navdrawerContainer, querySelector, toggleMenu;

    closeMenu = function() {
      body.classList.remove('open');
      appbarElement.classList.remove('open');
      navdrawerContainer.classList.remove('open');
    };
    toggleMenu = function() {
      body.classList.toggle('open');
      appbarElement.classList.toggle('open');
      navdrawerContainer.classList.toggle('open');
      navdrawerContainer.classList.add('opened');
    };
    'use strict';
    querySelector = document.querySelector.bind(document);
    navdrawerContainer = querySelector('.navdrawer-container');
    body = document.body;
    appbarElement = querySelector('.app-bar');
    menuBtn = querySelector('.menu');
    main = querySelector('main.8bd37191');
    main.addEventListener('click', closeMenu);
    menuBtn.addEventListener('click', toggleMenu);
    navdrawerContainer.addEventListener('click', function(event) {
      if (event.target.nodeName === 'A' || event.target.nodeName === 'LI') {
        closeMenu();
      }
    });
  })();

}).call(this);
