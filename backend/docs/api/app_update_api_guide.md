# 软件更新 API 测试指南

## 环境准备

确保已启动服务：
```bash
cd backend
make run
```

服务地址：`http://localhost:8080`

---

## 1. 检查更新

**GET** `/api/v1/app-updates/check`

检查指定平台和版本是否有新版本。

```bash
curl -X GET "http://localhost:8080/api/v1/app-updates/check?platform=android&version_code=1"
```

**请求参数：**
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| platform | string | 是 | 平台：`android`, `ios`, `web` |
| version_code | int | 是 | 当前版本代码 |

**响应示例（无更新）：**
```json
{
  "has_update": false,
  "force_update": false
}
```

**响应示例（有更新）：**
```json
{
  "has_update": true,
  "latest": {
    "id": 1,
    "version_code": 2,
    "version_name": "1.0.1",
    "platform": "android",
    "update_type": "minor",
    "is_force_update": false,
    "min_supported_version": "",
    "title": "体验优化更新",
    "description": "修复了一些已知问题，提升了使用体验。",
    "changelog": {
      "new_features": [
        "新增软件更新功能"
      ],
      "bug_fixes": [
        "修复首页显示问题"
      ]
    },
    "download_url": "/uploads/apk/android/2_app-release.apk",
    "file_size": 102400,
    "file_hash": "",
    "release_notes_url": "",
    "release_date": "2025-12-05T16:00:00Z"
  },
  "force_update": false,
  "update_reason": "发现新版本，建议更新"
}
```

---

## 2. 获取最新版本

**GET** `/api/v1/app-updates/latest`

获取指定平台的最新版本详情。

```bash
curl -X GET "http://localhost:8080/api/v1/app-updates/latest?platform=android"
```

**请求参数：**
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| platform | string | 是 | 平台：`android`, `ios`, `web` |

**响应示例：**
同"检查更新"中的 `latest` 对象。

---

## 3. 获取更新历史

**GET** `/api/v1/app-updates/history`

获取指定平台的更新历史记录。

```bash
curl -X GET "http://localhost:8080/api/v1/app-updates/history?platform=android"
```

**请求参数：**
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| platform | string | 是 | 平台：`android`, `ios`, `web` |

**响应示例：**
```json
[
  {
    "id": 2,
    "version_code": 2,
    "version_name": "1.0.1",
    ...
  },
  {
    "id": 1,
    "version_code": 1,
    "version_name": "1.0.0",
    ...
  }
]
```

---

## 4. 上传更新包

**POST** `/api/v1/app-updates`

上传APK文件并创建更新记录。

```bash
curl -X POST http://localhost:8080/api/v1/app-updates \
  -F "file=@/path/to/app-release.apk" \
  -F "version_code=2" \
  -F "version_name=1.0.1" \
  -F "platform=android" \
  -F "update_type=minor" \
  -F "title=体验优化更新" \
  -F "description=修复了一些已知问题" \
  -F "changelog={}"
```

**请求参数（Multipart/Form-Data）：**
| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| file | file | 是 | APK文件 |
| version_code | int | 是 | 版本代码（必须大于当前最新版本） |
| version_name | string | 是 | 版本名称（如 1.0.1） |
| platform | string | 是 | 平台：`android` |
| update_type | string | 是 | 更新类型：`major`, `minor`, `patch` |
| title | string | 是 | 更新标题 |
| description | string | 是 | 更新描述 |
| changelog | json string | 否 | 变更日志 |
| is_force_update | bool | 否 | 是否强制更新 |

**响应示例：**
```json
{
  "message": "Upload successful",
  "data": {
    "id": 2,
    "version_code": 2,
    ...
  }
}
```
