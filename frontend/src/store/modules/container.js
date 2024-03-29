import axios from 'axios';

const state = {
	containers: [],
	container: JSON.parse(localStorage.getItem('container')) || {},
	defaultContainer: {
		id: 0,
		title: 'Default'
	}
};

const mutations = {
	'SET_CONTAINERS'(state, containers) {
		state.containers = containers;
	},
	'SET_CONTAINER'(state, container) {
		state.container = container;
	}
};

const getters = {
	getContainer: state => state.container,
	getContainers: state => state.containers
};

const actions = {
	setDefaultContainer: ({commit, state}) => {
		commit('SET_CONTAINER', state.defaultContainer);
		localStorage.removeItem('container');
	},
	loadContainers: ({commit}) => {
		return new Promise((resolve, reject) => {
			axios.get("http://localhost:8888/api/containers")
			.then((res) => {
				commit('SET_CONTAINERS', res.data);
				resolve(res);
			})
			.catch((err) => {
				console.log(err);
				reject(err);
			});
		});
	},
	setContainer: ({commit}, container) => {
		commit('SET_CONTAINER', container);
		localStorage.setItem('container', container);
	},
	changeContainer: ({commit, state}, containerId) => {
		// if(containerId == null || containerId == '') {
		// 	commit('SET_CONTAINER', state.defaultContainer);
		// 	localStorage.setItem('container', JSON.stringify(state.defaultContainer));
		// 	return false;
		// }
		console.log('Change container: ', containerId)
		return new Promise((resolve, reject) => {
			axios.get('http://localhost:8888/api/containers/' + containerId)
			.then((res) => {
				console.log(res.data);
				commit('SET_CONTAINER', res.data);
				localStorage.setItem('container', JSON.stringify(res.data));
				resolve(res);
			})
			.catch((err) => {
				localStorage.removeItem('container');
				console.error(err);
				reject(err);
			});
		});
	},
	createContainer: ({commit}, data) => {
		return new Promise((resolve, reject) => {
			axios.post("http://localhost:8888/api/containers", {
				title: data.title
			})
			.then((res) => {
				resolve(res);
			})
			.catch((err) => {
				console.error(err);
				reject(err);
			});
		});
	},
	deleteContainer: ({commit}) => {
		return new Promise((resolve, reject) => {
			let container = JSON.parse(localStorage.getItem('container'));
			if(container.id != null) {
				axios.delete("http://localhost:8888/api/containers/" + container.id)
					.then((res) => {
						resolve(res);
					})
					.catch((err) => {
						console.error(err);
						reject(err);
					});
			} else {
				console.log("Container is not selected")
			}
		});
	},
	updateContainer: ({commit}, container) => {
		return new Promise((resolve, reject) => {
			axios.put('http://localhost:8888/api/containers/' + container.id, {
				title: container.title
			})
			.then((res) => {
				resolve(res);
			})
			.catch((err) => {
				console.error(err);
				reject(err);
			});
		});
	}
};

export default {
	state,
	mutations,
	getters,
	actions
};
