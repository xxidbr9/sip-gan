declare module namespace {

  export interface Resolution {
    128: string;
    640: string;
    1080: string;
  }

  interface Enhanced {
    full: Resolution;
    faces: Resolution[];
  }

  interface ImageAssets {
    originals: Resolution;
    enhanced: Enhanced;
  }

  export interface Gan {
    id: string;
    image_name: string;
    image_assets: ImageAssets;
    timestamp: string;
  }

  interface DataList {
    list_gan: Gan[];
  }

  export interface GanListResp {
    msg?: string;
    error?: Error;
    data: DataList;
  }

  export interface GanInfoResp {
    msg?: string;
    error?: Error;
    data: Gan;
  }

}

