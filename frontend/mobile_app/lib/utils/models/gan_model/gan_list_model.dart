class GanListModel {
  Null msg;
  Null error;
  Data data;

  GanListModel({this.msg, this.error, this.data});

  GanListModel.fromJson(Map<String, dynamic> json) {
    msg = json['msg'];
    error = json['error'];
    data = json['data'] != null ? new Data.fromJson(json['data']) : null;
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['msg'] = this.msg;
    data['error'] = this.error;
    if (this.data != null) {
      data['data'] = this.data.toJson();
    }
    return data;
  }
}

class Data {
  List<ListGan> listGan;

  Data({this.listGan});

  Data.fromJson(Map<String, dynamic> json) {
    if (json['list_gan'] != null) {
      listGan = new List<ListGan>();
      json['list_gan'].forEach((v) {
        listGan.add(new ListGan.fromJson(v));
      });
    }
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    if (this.listGan != null) {
      data['list_gan'] = this.listGan.map((v) => v.toJson()).toList();
    }
    return data;
  }
}

class ListGan {
  String id;
  String imageName;
  ImageAssets imageAssets;
  String timestamp;

  ListGan({this.id, this.imageName, this.imageAssets, this.timestamp});

  ListGan.fromJson(Map<String, dynamic> json) {
    id = json['id'];
    imageName = json['image_name'];
    imageAssets = json['image_assets'] != null
        ? new ImageAssets.fromJson(json['image_assets'])
        : null;
    timestamp = json['timestamp'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['id'] = this.id;
    data['image_name'] = this.imageName;
    if (this.imageAssets != null) {
      data['image_assets'] = this.imageAssets.toJson();
    }
    data['timestamp'] = this.timestamp;
    return data;
  }
}

class ImageAssets {
  Originals originals;
  Enhanced enhanced;

  ImageAssets({this.originals, this.enhanced});

  ImageAssets.fromJson(Map<String, dynamic> json) {
    originals = json['originals'] != null
        ? new Originals.fromJson(json['originals'])
        : null;
    enhanced = json['enhanced'] != null
        ? new Enhanced.fromJson(json['enhanced'])
        : null;
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    if (this.originals != null) {
      data['originals'] = this.originals.toJson();
    }
    if (this.enhanced != null) {
      data['enhanced'] = this.enhanced.toJson();
    }
    return data;
  }
}

class Originals {
  String s128;
  String s640;
  String s1080;

  Originals({this.s128, this.s640, this.s1080});

  Originals.fromJson(Map<String, dynamic> json) {
    s128 = json['128'];
    s640 = json['640'];
    s1080 = json['1080'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['128'] = this.s128;
    data['640'] = this.s640;
    data['1080'] = this.s1080;
    return data;
  }
}

class Faces {
  String s128;
  String s640;
  String s1080;

  Faces({this.s128, this.s640, this.s1080});

  Faces.fromJson(Map<String, dynamic> json) {
    s128 = json['128'];
    s640 = json['640'];
    s1080 = json['1080'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['128'] = this.s128;
    data['640'] = this.s640;
    data['1080'] = this.s1080;
    return data;
  }
}

class Enhanced {
  Originals full;
  List<Faces> faces;

  Enhanced({this.full, this.faces});

  Enhanced.fromJson(Map<String, dynamic> json) {
    full = json['full'] != null ? new Originals.fromJson(json['full']) : null;
    if (json['faces'] != null) {
      // ignore: deprecated_member_use
      faces = new List<Faces>();
      json['faces'].forEach((v) {
        faces.add(new Faces.fromJson(v));
      });
    }
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    if (this.full != null) {
      data['full'] = this.full.toJson();
    }
    if (this.faces != null) {
      data['faces'] = this.faces.map((v) => v.toJson()).toList();
    }
    return data;
  }
}
