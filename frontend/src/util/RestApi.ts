import Axios from 'axios';

const API_URL = process.env.VUE_APP_API_URL || `${window.location.origin}/api`;

export const RestApi = {
  file: {
    async getList(path: string) {
      return (await Axios.get(`${API_URL}/file/list?path=${path}`)).data.response;
    },
  },
};
