import axios from 'axios';
import qs from 'qs';

const state = {
	notes: [],
	note: {}
};

const mutations = {
	"SET_NOTES" (state, notes) {
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
	loadNotes: ({commit}, tab_id) => {
		return new Promise((resolve, reject) => {
			axios.get('http://localhost:8888/api/v1/notes/' + tab_id)
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
	loadNote: ({commit}, note_id) => {
		return new Promise((resolve, reject) => {
			axios.get(`http://localhost:8888/api/v1/notes/${note_id}/show`)
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
			axios.post('http://localhost:8888/api/v1/notes', {
				// data: qs.stringify(data)
				title: data.title,
				content: data.content,
				tab_id: getters.getCurrentTab.id,
				is_textual: data.is_textual
			})
			.then((res) => {
				resolve(res);
			})
			.catch((err) => {
				reject(err);
			});
		});
	},
	destroyNote: ({commit}, note_id) => {

		return new Promise((resolve, reject) => {
			axios.get(`http://localhost:8888/api/v1/notes/${note_id}/delete`)
			.then((res) => {
				resolve(res);
			})
			.catch((err) => {
				console.log(err);
				reject(err);
			});
		});
	},
	updateNote: ({commit}, data) => {
		return new Promise((resolve, reject) => {
			axios.post(`http://localhost:8888/api/v1/notes/${data.id}/update`, {
				title: data.title,
				content: data.content,
				id: data.id,
				tab_id: data.tab_id,
				is_textual: data.is_textual
			})
			.then((res) => {
				resolve(res);
			})
			.catch((err) => {
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
