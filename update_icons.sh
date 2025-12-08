#!/bin/bash

SOURCE_IMG="/Users/qiuhaonan/.gemini/antigravity/brain/55c8d93a-8db2-449b-89db-14d614527137/float_island_icon_squircle_1764927541211.png"
RES_DIR="/Users/qiuhaonan/Developer/projects/Float/frontend/android/app/src/main/res"

# Function to resize
resize_icon() {
    local density=$1
    local size=$2
    local target_dir="$RES_DIR/mipmap-$density"
    
    echo "Processing $density ($size x $size)..."
    
    if [ -d "$target_dir" ]; then
        sips -z $size $size "$SOURCE_IMG" --out "$target_dir/ic_launcher.png"
        sips -z $size $size "$SOURCE_IMG" --out "$target_dir/ic_launcher_round.png"
        
        if [ -f "$target_dir/ic_launcher_foreground.png" ]; then
             sips -z $size $size "$SOURCE_IMG" --out "$target_dir/ic_launcher_foreground.png"
        fi
    else
        echo "Directory $target_dir does not exist, skipping."
    fi
}

resize_icon "mdpi" 48
resize_icon "hdpi" 72
resize_icon "xhdpi" 96
resize_icon "xxhdpi" 144
resize_icon "xxxhdpi" 192

echo "Icon update complete."
