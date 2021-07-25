import Axios from 'axios';

const API_URL = process.env.VUE_APP_API_URL || `${window.location.origin}/api`;

export const RestApi = {
  file: {
    async getList(path: string) {
      return (await Axios.get(`${API_URL}/file/list?path=${path}`)).data.response;
    },
    async getDirSize(path: string): Promise<number> {
      return (await Axios.get(`${API_URL}/file/dirSize?path=${path}`)).data.response;
    },
    async open(path: string) {
      return (await Axios.post(`${API_URL}/file/open?path=${path}`)).data.response;
    },
    async rename(from: string, to: string) {
      return (
        await Axios.post(`${API_URL}/file/rename`, {
          from,
          to,
        })
      ).data.response;
    },
    async createDir(path: string) {
      return (await Axios.post(`${API_URL}/file/dir?path=${path}`)).data.response;
    },
    async createFile(path: string) {
      return (await Axios.post(`${API_URL}/file/file?path=${path}`)).data.response;
    },
    async deleteFile(path: string) {
      return (await Axios.delete(`${API_URL}/file/file?path=${path}`)).data.response;
    },
    async deleteDir(path: string) {
      return (await Axios.delete(`${API_URL}/file/dir?path=${path}`)).data.response;
    },
    async uploadFile(path: string, file: File, onProgess: (s: any) => void): Promise<string> {
      const formData = new FormData();
      formData.append(`file`, file, file.name);

      return (
        await Axios.post(`${API_URL}/file/file?path=${path}`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
          onUploadProgress: onProgess,
        })
      ).data.response;
    },

    async getPath() {
      return (await Axios.get(`${API_URL}/file/path?path`)).data.response;
    },
    async setPath(path: string) {
      return (await Axios.post(`${API_URL}/file/path?path=${path}`)).data.response;
    },
  },
};
