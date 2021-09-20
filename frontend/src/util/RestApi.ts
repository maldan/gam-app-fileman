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
    async deleteAny(path: string) {
      return (await Axios.delete(`${API_URL}/file/any?path=${path}`)).data.response;
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

    async copy(from: string, to: string) {
      return (
        await Axios.post(`${API_URL}/file/copy`, {
          from,
          to,
        })
      ).data.response;
    },

    async move(from: string, to: string) {
      return (
        await Axios.post(`${API_URL}/file/move`, {
          from,
          to,
        })
      ).data.response;
    },

    async saveInfo(path: string, data: any) {
      return (await Axios.post(`${API_URL}/file/info`, { path, data: JSON.stringify(data) })).data
        .response;
    },

    async getInfo(path: string) {
      try {
        const d = (await Axios.get(`${API_URL}/file/info?path=${path}`)).data.response;
        return JSON.parse(d);
      } catch {
        return {};
      }
    },

    async getFullInfo(path: string): Promise<any> {
      const d = (await Axios.get(`${API_URL}/file/fullInfo?path=${path}`)).data.response;
      return d;
    },
  },
  download: {
    async getList() {
      return (await Axios.get(`${API_URL}/download/list`)).data.response;
    },
    async add(url: string, path: string) {
      return (
        await Axios.post(`${API_URL}/download`, {
          url,
          path,
        })
      ).data.response;
    },
  },
  disk: {
    async getUsage(): Promise<any> {
      return (await Axios.get(`${API_URL}/disk/usage`)).data.response;
    },
  },
};
