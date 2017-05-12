import Vue from 'vue';
import VueRouter from 'vue-router';
import VueResource from 'vue-resource';

import util from './util.js'

import App from './App'
Vue.use(VueResource);

const app = new Vue({
	el: '#app',
	render: elem => elem(App),
	methods: {
		logged_in() {
			return !!(util.get_cookie('auth-token'));
		}
	},
});
