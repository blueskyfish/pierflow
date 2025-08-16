export enum RouteName {
  Projects = 'projects',
}

export enum RouteParam {
  ProjectId = 'projectId',
}

export class RoutePath {
  static HomePath = '/';
  static ProjectListPath = `/${RouteName.Projects}`;
  static ProjectHomePath = `/${RouteName.Projects}/:${RouteParam.ProjectId}`;
  static ProjectCommandPath = (commandPath: ProjectPath) =>
    `/${RouteName.Projects}/:${RouteParam.ProjectId}/${commandPath}`;
}

export enum ProjectPath {
  Detail = 'detail',
  Clone = 'clone',
  Checkout = 'checkout',
  Build = 'build',
  Start = 'start',
  Stop = 'stop',
  Pull = 'pull',
  Delete = 'delete',
  Create = 'create',
}

export class RouteBuilder {
  static buildProjectHomePath(projectId: string): string {
    return RoutePath.ProjectHomePath.replace(`:${RouteParam.ProjectId}`, projectId);
  }

  static buildProjectCommandPath(projectId: string, command: ProjectPath): string {
    return `${this.buildProjectHomePath(projectId)}/${command}`;
  }
}
