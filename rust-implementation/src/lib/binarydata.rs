/// Info about a binary system:
/// * `name` -- Name of binary system.
/// * `period` -- Period of the orbit.
/// * `epoch_peri` -- Epoch of the perihelion.
/// * `long_peri` -- Longitude of the perihelion.
/// * `ecc` -- Eccentricity of the orbit.
/// * `axis` -- Semi-major axis of the orbit.
/// * `incl` -- Orbital inclination.
/// * `pa_node` -- Position angle of the ascending node.
pub struct BinaryInfo {
    pub name: String,
    pub period: f64,
    pub epoch_peri: f64,
    pub long_peri: f64,
    pub ecc: f64,
    pub axis: f64,
    pub incl: f64,
    pub pa_node: f64,
}

/// Retrieve info about a binary system.
///
/// ## Returns
/// * BinaryInfo structure.
/// * status
pub fn get_binary_info_vector(binary_name: String) -> (BinaryInfo, String) {
    let mut binary_info_vector: Vec<BinaryInfo> = Vec::new();

    binary_info_vector.push(BinaryInfo {
        name: "eta-Cor".to_string(),
        period: 41.623,
        epoch_peri: 1934.008,
        long_peri: 219.907,
        ecc: 0.2763,
        axis: 0.907,
        incl: 59.025,
        pa_node: 23.717,
    });

    for i in binary_info_vector {
        if i.name == binary_name {
            return (i, "OK".to_string());
        }
    }

    return (
        BinaryInfo {
            name: binary_name,
            period: 0.0,
            epoch_peri: 0.0,
            long_peri: 0.0,
            ecc: 0.0,
            axis: 0.0,
            incl: 0.0,
            pa_node: 0.0,
        },
        "NotFound".to_string(),
    );
}
