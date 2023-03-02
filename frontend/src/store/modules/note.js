import axios from 'axios';

const state = {
	notes: [],
	note: {}
};

const mutations = {
	'SET_NOTES' (state, notes) {
		state.notes = notes;
	},
	'SET_NOTE' (state, note) {
		state.note = note;
	}
};

const getters = {
	getNotes: state => state.notes,
	getNote: state => state.note
};

const actions = {
	loadNotes: ({commit}, tabId) => {
		return new Promise((resolve, reject) => {
			axios.get('http://localhost:8888/api/notes/' + tabId + '/notes')
				.then((res) => {
					commit('SET_NOTES', res.data);
					resolve(res);
				})
				.catch((err) => {
					console.log(err);
					reject(err);
				})
		});
	},
	loadNote: ({commit}, noteId) => {
		return new Promise((resolve, reject) => {
			axios.get('http://localhost:8888/api/notes/' + noteId)
				.then((res) => {
					commit('SET_NOTE', res.data);
					resolve(res);
				})
				.catch((err) => {
					console.log(err);
					reject(err);
				})
		})
	},
	createNote: ({commit, state, getters}, data) => {
		return new Promise((resolve, reject) => {
			axios.post('http://localhost:8888/api/notes', {
				title: data.title,
				content: data.content,
				tabId: getters.getCurrentTab.id
			})
				.then((res) => {
					resolve(res);
				})
				.catch((err) => {
					reject(err);
				});
		});
	},
	updateNote: ({commit}, data) => {
		return new Promise((resolve, reject) => {
			axios.put('http://localhost:8888/api/notes/' + data.id, {
				title: data.title,
				content: data.content,
				id: data.id,
				tabId: data.tabId
			})
				.then((res) => {
					resolve(res);
				})
				.catch((err) => {
					reject(err);
				})
		})
	},
	destroyNote: ({commit}, noteId) => {
		return new Promise((resolve, reject) => {
			axios.delete('http://localhost:8888/api/notes/' + noteId)
				.then((res) => {
					resolve(res);
				})
				.catch((err) => {
					console.log(err);
					reject(err);
				});
		});
	},
	moveTo: ({commit}, tabId) => {
		return new Promise((resolve, reject) => {
			axios.post('http://localhost:8888/api/notes/move-to' + noteId, {
				tabId
			})
				.then(res => {
					resolve(res)
				})
				.catch((err) => {
					console.log(err);
					reject(err);
				});
		})
	}

};

export default {
	state,
	mutations,
	getters,
	actions
};
