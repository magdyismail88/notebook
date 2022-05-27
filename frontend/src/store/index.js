import Vue from 'vue';
import Vuex from 'vuex';
import tab from './modules/tab';
import note from './modules/note';
import container from './modules/container';

Vue.use(Vuex);

export default new Vuex.Store({
	modules: {
		tab,
		note,
		container
	}
});