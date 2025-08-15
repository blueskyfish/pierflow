export enum RouteName {
  Projects = 'projects',
}

export enum RouteParam {
  ProjectId = 'projectId',
}

export enum RoutePath {
  HomePath = '/',
  ProjectListPath = `/${RouteName.Projects}`,
  ProjectDetailPath = `/${RouteName.Projects}/:${RouteParam.ProjectId}`,
}

export class RouteBuilder {
  static buildProjectDetailPath(projectId: string): string {
    return RoutePath.ProjectDetailPath.replace(`:${RouteParam.ProjectId}`, projectId);
  }
}
