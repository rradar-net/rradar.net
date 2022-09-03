/* eslint-disable */
export const protobufPackage = "proto";

export enum Status {
  Success = 0,
  Fail = 1,
  Error = 2,
  UNRECOGNIZED = -1,
}

export function statusFromJSON(object: any): Status {
  switch (object) {
    case 0:
    case "Success":
      return Status.Success;
    case 1:
    case "Fail":
      return Status.Fail;
    case 2:
    case "Error":
      return Status.Error;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Status.UNRECOGNIZED;
  }
}

export function statusToJSON(object: Status): string {
  switch (object) {
    case Status.Success:
      return "Success";
    case Status.Fail:
      return "Fail";
    case Status.Error:
      return "Error";
    case Status.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}
