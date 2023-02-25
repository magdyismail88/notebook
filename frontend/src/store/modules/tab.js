import axios from 'axios';

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
			const container = JSON.parse(localStorage.getItem("container"));
			axios.get('http://localhost:8888/api/tabs/' + container.id + '/tabs')
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
			const container = JSON.parse(localStorage.getItem("container"));
			axios.post('http://localhost:8888/api/tabs', {
				title: data.title,
				containerId: container.id
			})
				.then((res) => {
					resolve(res);
				})
				.catch((err) => {
					reject(err);
				})
		});
	},
	removeTab: ({commit}, tabId) => {
		return new Promise((resolve, reject) => {
			axios.delete('http://localhost:8888/api/tabs/' + tabId)
				.then((res) => {
					resolve(res);
				})
				.catch((err) => {
					console.log(err);
					reject(err);
				})
		})
	},
	updateTab: ({commit}, tab) => {
		return new Promise((resolve, reject) => {
			axios.put('http://localhost:8888/api/tabs/' + tab.id, {
				id: tab.id,
				title: tab.title
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
	loadCurrentTab: ({commit}, tabId) => {
		return new Promise((resolve, reject) => {
			axios.get('http://localhost:8888/api/tabs/' + tabId)
					.then((res) => {
						commit('SET_CURRENT_TAB', res.data);
						console.log('TAB', res.data);
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
