// @generated by protoc-gen-es v1.3.0 with parameter "target=ts,import_extension=none"
// @generated from file cards/wheel.proto (package cards, syntax proto3)
/* eslint-disable */
// @ts-nocheck
import { proto3 } from "@bufbuild/protobuf";
/**
 * 麻將幸運輪
 *
 * @generated from enum cards.MahjongWheel
 */
export var MahjongWheel;
(function (MahjongWheel) {
    /**
     * 未指定
     *
     * @generated from enum value: MAHJONG_WEEL_CARD_UNSPECIFIED = 0;
     */
    MahjongWheel[MahjongWheel["MAHJONG_WEEL_CARD_UNSPECIFIED"] = 0] = "MAHJONG_WEEL_CARD_UNSPECIFIED";
    /**
     * @generated from enum value: MAHJONG_WEEL_EAST = 1;
     */
    MahjongWheel[MahjongWheel["MAHJONG_WEEL_EAST"] = 1] = "MAHJONG_WEEL_EAST";
    /**
     * @generated from enum value: MAHJONG_WEEL_SOUTH = 2;
     */
    MahjongWheel[MahjongWheel["MAHJONG_WEEL_SOUTH"] = 2] = "MAHJONG_WEEL_SOUTH";
    /**
     * @generated from enum value: MAHJONG_WEEL_WEST = 3;
     */
    MahjongWheel[MahjongWheel["MAHJONG_WEEL_WEST"] = 3] = "MAHJONG_WEEL_WEST";
    /**
     * @generated from enum value: MAHJONG_WEEL_NORTH = 4;
     */
    MahjongWheel[MahjongWheel["MAHJONG_WEEL_NORTH"] = 4] = "MAHJONG_WEEL_NORTH";
    /**
     * @generated from enum value: MAHJONG_WEEL_WHITE_DRAGON = 5;
     */
    MahjongWheel[MahjongWheel["MAHJONG_WEEL_WHITE_DRAGON"] = 5] = "MAHJONG_WEEL_WHITE_DRAGON";
    /**
     * @generated from enum value: MAHJONG_WEEL_RED_DRAGON = 6;
     */
    MahjongWheel[MahjongWheel["MAHJONG_WEEL_RED_DRAGON"] = 6] = "MAHJONG_WEEL_RED_DRAGON";
    /**
     * @generated from enum value: MAHJONG_WEEL_GREEN_DRAGON = 7;
     */
    MahjongWheel[MahjongWheel["MAHJONG_WEEL_GREEN_DRAGON"] = 7] = "MAHJONG_WEEL_GREEN_DRAGON";
})(MahjongWheel || (MahjongWheel = {}));
// Retrieve enum metadata with: proto3.getEnumType(MahjongWheel)
proto3.util.setEnumType(MahjongWheel, "cards.MahjongWheel", [
    { no: 0, name: "MAHJONG_WEEL_CARD_UNSPECIFIED" },
    { no: 1, name: "MAHJONG_WEEL_EAST" },
    { no: 2, name: "MAHJONG_WEEL_SOUTH" },
    { no: 3, name: "MAHJONG_WEEL_WEST" },
    { no: 4, name: "MAHJONG_WEEL_NORTH" },
    { no: 5, name: "MAHJONG_WEEL_WHITE_DRAGON" },
    { no: 6, name: "MAHJONG_WEEL_RED_DRAGON" },
    { no: 7, name: "MAHJONG_WEEL_GREEN_DRAGON" },
]);