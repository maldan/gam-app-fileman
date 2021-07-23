export class Helper {
  static fixPath(path: string): string {
    return path.replace(/\/\//g, '/');
  }
}
