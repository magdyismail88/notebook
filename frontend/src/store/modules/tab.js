import axios from 'axios';
import qs from 'qs';

const state = {
	tabs: [],
	tab: JSON.parse(localStorage.getItem('currentTab')) || {}
};

const mutations = {
	"SET_TABS" (state, tabs) {
		state.tabs = tabs;
	},
	"SET_CURRENT_TAB" (state, tab) {
		state.tab = tab;
	}
};

const getters = {
	getTabs: state => state.tabs,
	getCurrentTab: state => state.tab
};

const actions = {
	loadTabs: ({commit}) => {
		return new Promise((resolve, reject) => {
			let container = JSON.parse(localStorage.getItem("container"));
			axios.get(`http://localhost:8888/api/tabs/${container.id}`)
			.then((res) => {
				commit('SET_TABS', res.data);
				resolve(res);
			})
			.catch((err) => {
				reject(err);
			})
		});
	},
	createTab: ({commit}, data) => {
		return new Promise((resolve, reject) => {

			let container = JSON.parse(localStorage.getItem("container"));

			axios.post("http://localhost:8888/api/tabs", {
				// headers: {
				// 	'content-type': 'application/x-www-form-urlencoded',
				// 	// 'content-type': 'application/json',
				// },
				title: data.title,
				container_id: container.id
			})
			.then((res) => {
				resolve(res);
			})
			.catch((err) => {
				reject(err);
			})
		});
	},
	removeTab: ({commit}, tab_id) => {
		return new Promise((resolve, reject) => {
			axios.get(`http://localhost:8888/api/tabs/${tab_id}/delete`)
			.then((res) => {
				resolve(res);
			})
			.catch((err) => {
				console.log(err);
				reject(err);
			})
		})
	},
	updateTab: ({commit}, data) => {
		return new Promise((resolve, reject) => {
			axios.post(`http://localhost:8888/api/tabs/${data.id}/update`, {
				id: data.id,
				title: data.title
			})
			.then((res) => {
				resolve(res);
			})
			.catch((err) => {
				console.log(err);
				reject(err);
			});
		});
	},
	loadCurrentTab: ({commit}, tab_id) => {
		return new Promise((resolve, reject) => {
			axios.get(`http://localhost:8888/api/tabs/${tab_id}/tab`)
				.then((res) => {
					commit('SET_CURRENT_TAB', res.data);
					localStorage.setItem('tab', JSON.stringify(res.data))
					resolve(res);
				})
				.catch((err) => {
					console.log(err);
					reject(err);
				})
		})
	}
};


export default {
	state,
	mutations,
	getters,
	actions
};
