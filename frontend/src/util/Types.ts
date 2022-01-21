export interface IFile {
  name: string;
  path: string;
  kind: 'dir' | 'file';
  size: number;
  user: string;
  tags: Record<string, string | number | boolean>;
  created: string;
}
