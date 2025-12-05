INSERT INTO app_updates (
    version_code, version_name, platform, update_type, is_force_update, 
    title, description, changelog, download_url, status, release_date
) VALUES (
    2, '1.0.1', 'android', 'minor', false, 
    '体验优化更新', '修复了一些已知问题，提升了使用体验。', 
    '{"new_features": ["新增软件更新功能"], "bug_fixes": ["修复首页显示问题"]}', 
    'https://example.com/float-android-1.0.1.apk', 'released', NOW()
);
